//go:generate goversioninfo -manifest=testdata/resource/goversioninfo.exe.manifest
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/pedroalbanese/kp/internal/commands"
	t "github.com/pedroalbanese/kp/internal/backend/types"
	v1 "github.com/pedroalbanese/kp/internal/backend/keepassv1"
)

var (
	keyFile        = flag.String("key", "", "a key file to use to unlock the db")
	dbFile         = flag.String("db", "", "the db to open")
	noninteractive = flag.String("n", "", "execute a given command and exit")
	version        = flag.Bool("version", false, "print version and exit")
)

const (
	buildVersionString = "0.0.2"
)

// promptForDBPassword will determine the password based on environment vars or, lacking those, a prompt to the user
func promptForDBPassword(shell *ishell.Shell) (string) {
	// we are prompting for the password
	shell.Print("enter database password: ")
	return shell.ReadPassword()
}

// newDB will create or open a DB with the parameters specified.  `open` indicates whether the DB should be opened or not (vs created)
func newDB(dbPath string, password string, keyPath string, version int) (t.Database, error) {
	var dbWrapper t.Database
	dbWrapper = &v1.Database{}
	dbOpts := t.Options {
		DBPath: dbPath,
		Password: password,
		KeyPath: keyPath,
	}
	err := dbWrapper.Init(dbOpts)
	return dbWrapper, err
}

func main() {
	flag.Parse()

	shell := ishell.New()
	if *version {
		shell.Printf("version: %s\n", buildVersionString)
		os.Exit(1)
	}

	var dbWrapper t.Database

	dbPath, exists := os.LookupEnv("KP_DATABASE")
	if !exists {
		dbPath = *dbFile
	}

	// default to the flag argument
	keyPath := *keyFile

	if envKeyfile, found := os.LookupEnv("KP_KEYFILE"); found && keyPath == "" {
		keyPath = envKeyfile
	}

	for {
		// if the password is coming from an environment variable, we need to terminate
		// after the first attempt or it will fall into an infinite loop
		var err error
		password, passwordInEnv := os.LookupEnv("KP_PASSWORD")
		if !passwordInEnv {
			password = promptForDBPassword(shell)

			if err != nil {
				shell.Printf("could not retrieve password: %s", err)
				os.Exit(1)
			}
		}

		dbWrapper, err = newDB(dbPath, password, keyPath, 1)
		if err != nil {
			// typically, these errors will be a bad password, so we want to keep prompting until the user gives up
			// if, however, the password is in an environment variable, we want to abort immediately so the program doesn't fall
			// in to an infinite loop
			shell.Printf("could not open database: %s\n", err)
			if passwordInEnv {
				os.Exit(1)
			}
			continue
		}
		break
	}

//	if dbWrapper.Locked() {
//		shell.Printf("Lockfile exists for DB at path '%s', another process is using this database!\n", dbWrapper.SavePath())
//		shell.Printf("Open anyways? Data loss may occur. (will only proceed if 'yes' is entered)  ")
//		line, err := shell.ReadLineErr()
//		if err != nil {
//			shell.Printf("could not read user input: %s\n", line)
//			os.Exit(1)
//		}

//		if line != "yes" {
//			shell.Println("aborting")
//			os.Exit(1)
//		}
//	}

//	if err := dbWrapper.Lock(); err != nil {
//		shell.Printf("aborting, could not lock database: %s\n", err)
//		os.Exit(1)
//	}

//	shell.Printf("opened database at %s\n", dbWrapper.SavePath())

	shell.Set("db", dbWrapper)
	shell.SetPrompt(fmt.Sprintf("/%s > ", dbWrapper.CurrentLocation().Name()))

	shell.AddCmd(&ishell.Cmd{
		Name:                "ls",
		Help:                "ls [path]",
		Func:                commands.Ls(shell),
	})
	shell.AddCmd(&ishell.Cmd{
		Name:                "new",
		Help:                "new <path>",
		LongHelp:            "creates a new entry at <path>",
		Func:                commands.NewEntry(shell),
	})
	shell.AddCmd(&ishell.Cmd{
		Name:                "mkdir",
		LongHelp:            "create a new group",
		Help:                "mkdir <group name>",
		Func:                commands.NewGroup(shell),
	})
	shell.AddCmd(&ishell.Cmd{
		Name:     "saveas",
		LongHelp: "save this db to a new file, existing credentials will be used, the new location will /not/ be used as the main save path",
		Help:     "saveas <file.kdb>",
		Func:     commands.SaveAs(shell),
	})

	if dbWrapper.Version() == t.V2 {
		shell.AddCmd(&ishell.Cmd{
			Name:                "select",
			Help:                "select [-f] <entry>",
			LongHelp:            "shows details on a given value in an entry, passwords will be redacted unless '-f' is specified",
			Func:                commands.Select(shell),
		})
	}

	shell.AddCmd(&ishell.Cmd{
		Name:                "show",
		Help:                "show [-f] <entry>",
		LongHelp:            "shows details on a given entry, passwords will be redacted unless '-f' is specified",
		Func:                commands.Show(shell),
	})
	shell.AddCmd(&ishell.Cmd{
		Name:                "cd",
		Help:                "cd <path>",
		LongHelp:            "changes the current group to a different path",
		Func:                commands.Cd(shell),
	})

	attachCmd := &ishell.Cmd{
		Name:     "attach",
		LongHelp: "manages the attachment for a given entry",
		Help:     "attach <get|show|delete> <entry> <filesystem location>",
	}
	attachCmd.AddCmd(&ishell.Cmd{
		Name:                "create",
		Help:                "attach create <entry> <name> <filesystem location>",
		LongHelp:            "creates a new attachment based on a local file",
		Func:                commands.Attach(shell, "create"),
	})
	attachCmd.AddCmd(&ishell.Cmd{
		Name:                "get",
		Help:                "attach get <entry> <filesystem location> // ('-' for stdout)",
		LongHelp:            "retrieves an attachment and outputs it to a filesystem location",
		Func:                commands.Attach(shell, "get"),
	})
	attachCmd.AddCmd(&ishell.Cmd{
		Name:                "details",
		Help:                "attach details <entry>",
		LongHelp:            "shows the details of the attachment on an entry",
		Func:                commands.Attach(shell, "details"),
	})
	shell.AddCmd(attachCmd)

	shell.AddCmd(&ishell.Cmd{
		LongHelp:            "searches for any entries with the regular expression '<term>' in their titles or contents",
		Name:                "search",
		Help:                "search <term>",
		Func:                commands.Search(shell),
	})

	shell.AddCmd(&ishell.Cmd{
		Name:                "rm",
		Help:                "rm <entry>",
		LongHelp:            "removes an entry",
		Func:                commands.Rm(shell),
	})

	shell.AddCmd(&ishell.Cmd{
		Name:                "xp",
		Help:                "xp <entry>",
		LongHelp:            "copies a password to the clipboard",
		Func:                commands.Xp(shell),
	})

	shell.AddCmd(&ishell.Cmd{
		Name:                "xk",
		Help:                "xk <entry>",
		LongHelp:            "copies a password to the clipboard",
		Func:                commands.Xk(shell),
	})

	shell.AddCmd(&ishell.Cmd{
		Name:                "edit",
		Help:                "edit <entry>",
		LongHelp:            "edits an existing entry",
		Func:                commands.Edit(shell),
	})

	shell.AddCmd(&ishell.Cmd{
		Name:     "pwd",
		Help:     "pwd",
		LongHelp: "shows path of current group",
		Func:     commands.Pwd(shell),
	})

	shell.AddCmd(&ishell.Cmd{
		Name:     "save",
		Help:     "save",
		LongHelp: "saves the database to its most recently used path",
		Func:     commands.Save(shell),
	})

	shell.AddCmd(&ishell.Cmd{
		Name:     "xx",
		Help:     "xx",
		LongHelp: "clears the clipboard",
		Func:     commands.Xx(shell),
	})

	shell.AddCmd(&ishell.Cmd{
		Name:                "xu",
		Help:                "xu",
		LongHelp:            "copies username to the clipboard",
		Func:                commands.Xu(shell),
	})

	shell.AddCmd(&ishell.Cmd{
		Name:                "xw",
		Help:                "xw",
		LongHelp:            "copies url to clipboard",
		Func:                commands.Xw(shell),
	})

	shell.AddCmd(&ishell.Cmd{
		Name:                "mv",
		Help:                "mv <source> <destination>",
		LongHelp:            "moves entries between groups",
		Func:                commands.Mv(shell),
	})

	shell.AddCmd(&ishell.Cmd{
		Name:     "version",
		Help:     "version",
		LongHelp: "prints version",
		Func: func(c *ishell.Context) {
			shell.Printf("version: %s\n", buildVersionString)
		},
	})

	if *noninteractive != "" {
		bits := strings.Split(*noninteractive, " ")
		if err := shell.Process([]string{bits[0], strings.Join(bits[1:], " ")}...); err != nil {
			shell.Printf("error processing command: %s\n", err)
		}
	} else {
		shell.Run()
	}

	// This will run after the shell exits
	fmt.Fprintln(os.Stderr, "exiting")

	if dbWrapper.Changed() {
		if err := commands.PromptAndSave(shell); err != nil {
			fmt.Printf("error attempting to save database: %s\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "no changes detected since last save.")
	}


//	if err := dbWrapper.Unlock(); err != nil {
//		fmt.Printf("failed to unlock db: %s", err)
//		os.Exit(1)
//	}
}
