# golib - Library of Go packages

[//]: # (GODOC)
- [appdata](#github.com/billziss-gh/golib/appdata) - Package appdata provides access to well known directories for applications.
- [cmd](#github.com/billziss-gh/golib/cmd) - Package cmd provides (sub-)command functionality for command-line programs.
- [config](#github.com/billziss-gh/golib/config) - Package config is used to read and write configuration files.
- [errors](#github.com/billziss-gh/golib/errors) - Package errors implements functions for advanced error handling.
- [keyring](#github.com/billziss-gh/golib/keyring) - Package keyring implements functions for accessing and storing passwords in the system's keyring (Keychain on macOS, Credential Manager on Windows, Secret Service on Linux).
- [retry](#github.com/billziss-gh/golib/retry) - Package retry implements simple retry functionality.
- [trace](#github.com/billziss-gh/golib/trace) - Package trace provides a simple tracing facility for Go functions.
- [util](#github.com/billziss-gh/golib/util) - Package util contains general utility functions.



----
## <a name="github.com/billziss-gh/golib/appdata">Package appdata</a>
_[[godoc.org](https://godoc.org/github.com/billziss-gh/golib/appdata)]_

`import "github.com/billziss-gh/golib/appdata"`

* [Overview](#github.com/billziss-gh/golib/appdata/pkg-overview)
* [Index](#github.com/billziss-gh/golib/appdata/pkg-index)

### <a name="github.com/billziss-gh/golib/appdata/pkg-overview">Overview</a>
Package appdata provides access to well known directories for applications.




### <a name="github.com/billziss-gh/golib/appdata/pkg-index">Index</a>
* [Constants](#github.com/billziss-gh/golib/appdata/pkg-constants)
* [func CacheDir() (string, error)](#github.com/billziss-gh/golib/appdata/CacheDir)
* [func ConfigDir() (string, error)](#github.com/billziss-gh/golib/appdata/ConfigDir)
* [func DataDir() (string, error)](#github.com/billziss-gh/golib/appdata/DataDir)
* [type AppData](#github.com/billziss-gh/golib/appdata/AppData)


##### <a name="github.com/billziss-gh/golib/appdata/pkg-files">Package files</a>
[appdata.go](appdata/appdata.go) [appdata_darwin.go](appdata/appdata_darwin.go) 


### <a name="github.com/billziss-gh/golib/appdata/pkg-constants">Constants</a>
``` go
const ErrAppData = "ErrAppData"
```



### <a name="github.com/billziss-gh/golib/appdata/CacheDir">func</a> [CacheDir](appdata/appdata.go#L42)
``` go
func CacheDir() (string, error)
```
CacheDir returns the directory where application cache files
should be stored.



### <a name="github.com/billziss-gh/golib/appdata/ConfigDir">func</a> [ConfigDir](appdata/appdata.go#L30)
``` go
func ConfigDir() (string, error)
```
ConfigDir returns the directory where application configuration files
should be stored.



### <a name="github.com/billziss-gh/golib/appdata/DataDir">func</a> [DataDir](appdata/appdata.go#L36)
``` go
func DataDir() (string, error)
```
DataDir returns the directory where application data files
should be stored.




### <a name="github.com/billziss-gh/golib/appdata/AppData">type</a> [AppData](appdata/appdata.go#L20)
``` go
type AppData interface {
    ConfigDir() (string, error)
    DataDir() (string, error)
    CacheDir() (string, error)
}
```

``` go
var DefaultAppData AppData
```
















----
## <a name="github.com/billziss-gh/golib/cmd">Package cmd</a>
_[[godoc.org](https://godoc.org/github.com/billziss-gh/golib/cmd)]_

`import "github.com/billziss-gh/golib/cmd"`

* [Overview](#github.com/billziss-gh/golib/cmd/pkg-overview)
* [Index](#github.com/billziss-gh/golib/cmd/pkg-index)

### <a name="github.com/billziss-gh/golib/cmd/pkg-overview">Overview</a>
Package cmd provides (sub-)command functionality for command-line programs.
This package works closely with the standard library flag package.




### <a name="github.com/billziss-gh/golib/cmd/pkg-index">Index</a>
* [Variables](#github.com/billziss-gh/golib/cmd/pkg-variables)
* [func PrintCmds()](#github.com/billziss-gh/golib/cmd/PrintCmds)
* [func Run()](#github.com/billziss-gh/golib/cmd/Run)
* [type Cmd](#github.com/billziss-gh/golib/cmd/Cmd)
  * [func Add(name string, main func(*Cmd, []string)) *Cmd](#github.com/billziss-gh/golib/cmd/Add)
  * [func (self *Cmd) GetFlag(name string) interface{}](#github.com/billziss-gh/golib/cmd/Cmd.GetFlag)
* [type CmdMap](#github.com/billziss-gh/golib/cmd/CmdMap)
  * [func NewCmdMap() *CmdMap](#github.com/billziss-gh/golib/cmd/NewCmdMap)
  * [func (self *CmdMap) Add(name string, main func(*Cmd, []string)) (cmd *Cmd)](#github.com/billziss-gh/golib/cmd/CmdMap.Add)
  * [func (self *CmdMap) Get(name string) *Cmd](#github.com/billziss-gh/golib/cmd/CmdMap.Get)
  * [func (self *CmdMap) GetNames() []string](#github.com/billziss-gh/golib/cmd/CmdMap.GetNames)
  * [func (self *CmdMap) PrintCmds()](#github.com/billziss-gh/golib/cmd/CmdMap.PrintCmds)


##### <a name="github.com/billziss-gh/golib/cmd/pkg-files">Package files</a>
[cmd.go](cmd/cmd.go) 



### <a name="github.com/billziss-gh/golib/cmd/pkg-variables">Variables</a>
``` go
var DefaultCmdMap = NewCmdMap()
```
DefaultCmdMap is the default command map.



### <a name="github.com/billziss-gh/golib/cmd/PrintCmds">func</a> [PrintCmds](cmd/cmd.go#L141)
``` go
func PrintCmds()
```
PrintCmds prints help text for all commands in the default command map
to stderr.



### <a name="github.com/billziss-gh/golib/cmd/Run">func</a> [Run](cmd/cmd.go#L146)
``` go
func Run()
```
Run parses the command line and executes the specified (sub-)command.




### <a name="github.com/billziss-gh/golib/cmd/Cmd">type</a> [Cmd](cmd/cmd.go#L33)
``` go
type Cmd struct {
    // Flag contains the command flag set.
    Flag *flag.FlagSet

    // Main is the function to run when the command is selected.
    Main func(cmd *Cmd, args []string)

    // Use contains the command usage string.
    Use string

    // Desc contains the command description.
    Desc string
}
```
Cmd encapsulates a (sub-)command.







#### <a name="github.com/billziss-gh/golib/cmd/Add">func</a> [Add](cmd/cmd.go#L135)
``` go
func Add(name string, main func(*Cmd, []string)) *Cmd
```
Add adds a new command in the default command map.

The name parameter is the command name. However if this parameter contains
a space or newline it is interpreted as described below. Consider:


	NAME ARGUMENTS
	DESCRIPTION

Then the command name becomes "NAME", the command Use field becomes
"NAME ARGUMENTS" and the command Desc field becomes "DESCRIPTION".





#### <a name="github.com/billziss-gh/golib/cmd/Cmd.GetFlag">func</a> (\*Cmd) [GetFlag](cmd/cmd.go#L113)
``` go
func (self *Cmd) GetFlag(name string) interface{}
```
GetFlag gets the value of the named flag.




### <a name="github.com/billziss-gh/golib/cmd/CmdMap">type</a> [CmdMap](cmd/cmd.go#L26)
``` go
type CmdMap struct {
    // contains filtered or unexported fields
}
```
CmdMap encapsulates a (sub-)command map.







#### <a name="github.com/billziss-gh/golib/cmd/NewCmdMap">func</a> [NewCmdMap](cmd/cmd.go#L106)
``` go
func NewCmdMap() *CmdMap
```
NewCmdMap creates a new command map.





#### <a name="github.com/billziss-gh/golib/cmd/CmdMap.Add">func</a> (\*CmdMap) [Add](cmd/cmd.go#L57)
``` go
func (self *CmdMap) Add(name string, main func(*Cmd, []string)) (cmd *Cmd)
```
Add adds a new command in the command map.

The name parameter is the command name. However if this parameter contains
a space or newline it is interpreted as described below. Consider:


	NAME ARGUMENTS
	DESCRIPTION

Then the command name becomes "NAME", the command Use field becomes
"NAME ARGUMENTS" and the command Desc field becomes "DESCRIPTION".




#### <a name="github.com/billziss-gh/golib/cmd/CmdMap.Get">func</a> (\*CmdMap) [Get](cmd/cmd.go#L76)
``` go
func (self *CmdMap) Get(name string) *Cmd
```
Get gets a command by name.




#### <a name="github.com/billziss-gh/golib/cmd/CmdMap.GetNames">func</a> (\*CmdMap) [GetNames](cmd/cmd.go#L83)
``` go
func (self *CmdMap) GetNames() []string
```
GetNames gets all command names.




#### <a name="github.com/billziss-gh/golib/cmd/CmdMap.PrintCmds">func</a> (\*CmdMap) [PrintCmds](cmd/cmd.go#L92)
``` go
func (self *CmdMap) PrintCmds()
```
PrintCmds prints help text for all commands to stderr.











----
## <a name="github.com/billziss-gh/golib/config">Package config</a>
_[[godoc.org](https://godoc.org/github.com/billziss-gh/golib/config)]_

`import "github.com/billziss-gh/golib/config"`

* [Overview](#github.com/billziss-gh/golib/config/pkg-overview)
* [Index](#github.com/billziss-gh/golib/config/pkg-index)

### <a name="github.com/billziss-gh/golib/config/pkg-overview">Overview</a>
Package config is used to read and write configuration files.

Configuration files are similar to Windows INI files. They store a list
of properties (key/value pairs); they may also be grouped into sections.

The basic syntax of a configuration file is as follows:


	name1=value1
	name2=value2
	...
	
	[section]
	name3=value3
	name4=value4
	...

Properties not in a section are placed in the unnamed (empty "") section.




### <a name="github.com/billziss-gh/golib/config/pkg-index">Index</a>
* [Variables](#github.com/billziss-gh/golib/config/pkg-variables)
* [func ReadFunc(reader io.Reader, fn func(sect, name string, valu interface{})) error](#github.com/billziss-gh/golib/config/ReadFunc)
* [func Write(writer io.Writer, conf Config) error](#github.com/billziss-gh/golib/config/Write)
* [func WriteTyped(writer io.Writer, conf TypedConfig) error](#github.com/billziss-gh/golib/config/WriteTyped)
* [type Config](#github.com/billziss-gh/golib/config/Config)
  * [func Read(reader io.Reader) (Config, error)](#github.com/billziss-gh/golib/config/Read)
* [type Dialect](#github.com/billziss-gh/golib/config/Dialect)
  * [func (dialect *Dialect) Read(reader io.Reader) (Config, error)](#github.com/billziss-gh/golib/config/Dialect.Read)
  * [func (dialect *Dialect) ReadFunc(reader io.Reader, fn func(sect, name string, valu interface{})) error](#github.com/billziss-gh/golib/config/Dialect.ReadFunc)
  * [func (dialect *Dialect) ReadTyped(reader io.Reader) (TypedConfig, error)](#github.com/billziss-gh/golib/config/Dialect.ReadTyped)
  * [func (dialect *Dialect) Write(writer io.Writer, conf Config) error](#github.com/billziss-gh/golib/config/Dialect.Write)
  * [func (dialect *Dialect) WriteTyped(writer io.Writer, conf TypedConfig) error](#github.com/billziss-gh/golib/config/Dialect.WriteTyped)
* [type Section](#github.com/billziss-gh/golib/config/Section)
* [type TypedConfig](#github.com/billziss-gh/golib/config/TypedConfig)
  * [func ReadTyped(reader io.Reader) (TypedConfig, error)](#github.com/billziss-gh/golib/config/ReadTyped)
* [type TypedSection](#github.com/billziss-gh/golib/config/TypedSection)


##### <a name="github.com/billziss-gh/golib/config/pkg-files">Package files</a>
[config.go](config/config.go) 



### <a name="github.com/billziss-gh/golib/config/pkg-variables">Variables</a>
``` go
var DefaultDialect = &Dialect{
    AssignChars:    "=:",
    CommentChars:   ";#",
    ReadEmptyKeys:  true,
    WriteEmptyKeys: false,
    Strict:         false,
}
```
DefaultDialect contains the default configuration dialect.
It is compatible with Windows INI files.



### <a name="github.com/billziss-gh/golib/config/ReadFunc">func</a> [ReadFunc](config/config.go#L424)
``` go
func ReadFunc(
    reader io.Reader, fn func(sect, name string, valu interface{})) error
```


### <a name="github.com/billziss-gh/golib/config/Write">func</a> [Write](config/config.go#L443)
``` go
func Write(writer io.Writer, conf Config) error
```
Write writes a configuration to the supplied writer
using the default dialect.



### <a name="github.com/billziss-gh/golib/config/WriteTyped">func</a> [WriteTyped](config/config.go#L449)
``` go
func WriteTyped(writer io.Writer, conf TypedConfig) error
```
WriteTyped writes a typed configuration to the supplied writer
using the default dialect.




### <a name="github.com/billziss-gh/golib/config/Config">type</a> [Config](config/config.go#L49)
``` go
type Config map[string]Section
```
Config is used to store a configuration as string properties.







#### <a name="github.com/billziss-gh/golib/config/Read">func</a> [Read](config/config.go#L431)
``` go
func Read(reader io.Reader) (Config, error)
```
Read reads a configuration from the supplied reader
using the default dialect.





### <a name="github.com/billziss-gh/golib/config/Dialect">type</a> [Dialect](config/config.go#L59)
``` go
type Dialect struct {
    // AssignChars contains the characters used for property assignment.
    // The first character in AssignChars is the character used during
    // writing.
    AssignChars string

    // CommentChars contains the characters used for comments.
    CommentChars string

    // ReadEmptyKeys determines whether to read properties with missing values.
    // The properties so created will be interpretted as empty strings for Read
    // and boolean true for ReadTyped.
    ReadEmptyKeys bool

    // WriteEmptyKeys determines whether to write properties with missing values.
    // This is only important when writing boolean true properties with
    // WriteTyped; these will be written with missing values.
    WriteEmptyKeys bool

    // Strict determines whether parse errors should be reported.
    Strict bool
}
```
Dialect is used to represent different dialects of configuration files.










#### <a name="github.com/billziss-gh/golib/config/Dialect.Read">func</a> (\*Dialect) [Read](config/config.go#L169)
``` go
func (dialect *Dialect) Read(reader io.Reader) (Config, error)
```
Read reads a configuration from the supplied reader.




#### <a name="github.com/billziss-gh/golib/config/Dialect.ReadFunc">func</a> (\*Dialect) [ReadFunc](config/config.go#L92)
``` go
func (dialect *Dialect) ReadFunc(
    reader io.Reader, fn func(sect, name string, valu interface{})) error
```



#### <a name="github.com/billziss-gh/golib/config/Dialect.ReadTyped">func</a> (\*Dialect) [ReadTyped](config/config.go#L193)
``` go
func (dialect *Dialect) ReadTyped(reader io.Reader) (TypedConfig, error)
```
ReadTyped reads a typed configuration from the supplied reader.




#### <a name="github.com/billziss-gh/golib/config/Dialect.Write">func</a> (\*Dialect) [Write](config/config.go#L239)
``` go
func (dialect *Dialect) Write(writer io.Writer, conf Config) error
```
Write writes a configuration to the supplied writer.




#### <a name="github.com/billziss-gh/golib/config/Dialect.WriteTyped">func</a> (\*Dialect) [WriteTyped](config/config.go#L280)
``` go
func (dialect *Dialect) WriteTyped(writer io.Writer, conf TypedConfig) error
```
WriteTyped writes a typed configuration to the supplied writer.




### <a name="github.com/billziss-gh/golib/config/Section">type</a> [Section](config/config.go#L46)
``` go
type Section map[string]string
```
Section is used to store a configuration section as string properties.










### <a name="github.com/billziss-gh/golib/config/TypedConfig">type</a> [TypedConfig](config/config.go#L55)
``` go
type TypedConfig map[string]TypedSection
```
TypedConfig is used to store a configuration as typed properties.







#### <a name="github.com/billziss-gh/golib/config/ReadTyped">func</a> [ReadTyped](config/config.go#L437)
``` go
func ReadTyped(reader io.Reader) (TypedConfig, error)
```
ReadTyped reads a typed configuration from the supplied reader
using the default dialect.





### <a name="github.com/billziss-gh/golib/config/TypedSection">type</a> [TypedSection](config/config.go#L52)
``` go
type TypedSection map[string]interface{}
```
TypedSection is used to store a configuration section as typed properties.

















----
## <a name="github.com/billziss-gh/golib/errors">Package errors</a>
_[[godoc.org](https://godoc.org/github.com/billziss-gh/golib/errors)]_

`import "github.com/billziss-gh/golib/errors"`

* [Overview](#github.com/billziss-gh/golib/errors/pkg-overview)
* [Index](#github.com/billziss-gh/golib/errors/pkg-index)

### <a name="github.com/billziss-gh/golib/errors/pkg-overview">Overview</a>
Package errors implements functions for advanced error handling.

Errors in this package contain a message, a cause (an error that caused
this error) and an attachment (any interface{}). Errors also contain
information about the program location where they were created.

Errors can be printed using the fmt.Printf verbs %s, %q, %x, %X, %v. In
particular the %+v format will print an error complete with its stack trace.

Inspired by <a href="https://github.com/pkg/errors">https://github.com/pkg/errors</a>




### <a name="github.com/billziss-gh/golib/errors/pkg-index">Index</a>
* [func Attachment(err error) interface{}](#github.com/billziss-gh/golib/errors/Attachment)
* [func Cause(err error) error](#github.com/billziss-gh/golib/errors/Cause)
* [func HasAttachment(err error, attachment interface{}) bool](#github.com/billziss-gh/golib/errors/HasAttachment)
* [func HasCause(err error, cause error) bool](#github.com/billziss-gh/golib/errors/HasCause)
* [func New(message string, args ...interface{}) error](#github.com/billziss-gh/golib/errors/New)


##### <a name="github.com/billziss-gh/golib/errors/pkg-files">Package files</a>
[errors.go](errors/errors.go) 





### <a name="github.com/billziss-gh/golib/errors/Attachment">func</a> [Attachment](errors/errors.go#L185)
``` go
func Attachment(err error) interface{}
```
Attachment will return additional information attached to this error
(if any).



### <a name="github.com/billziss-gh/golib/errors/Cause">func</a> [Cause](errors/errors.go#L174)
``` go
func Cause(err error) error
```
Cause will return the error that caused this error (if any).



### <a name="github.com/billziss-gh/golib/errors/HasAttachment">func</a> [HasAttachment](errors/errors.go#L207)
``` go
func HasAttachment(err error, attachment interface{}) bool
```
HasAttachment determines if a particular attachment is in the causal chain
of this error.



### <a name="github.com/billziss-gh/golib/errors/HasCause">func</a> [HasCause](errors/errors.go#L196)
``` go
func HasCause(err error, cause error) bool
```
HasCause determines if a particular error is in the causal chain
of this error.



### <a name="github.com/billziss-gh/golib/errors/New">func</a> [New](errors/errors.go#L158)
``` go
func New(message string, args ...interface{}) error
```
New creates an error with a message. Additionally the error may contain
a cause (an error that caused this error) and an attachment (any
interface{}). New will also record information about the program location
where it was called.











----
## <a name="github.com/billziss-gh/golib/keyring">Package keyring</a>
_[[godoc.org](https://godoc.org/github.com/billziss-gh/golib/keyring)]_

`import "github.com/billziss-gh/golib/keyring"`

* [Overview](#github.com/billziss-gh/golib/keyring/pkg-overview)
* [Index](#github.com/billziss-gh/golib/keyring/pkg-index)

### <a name="github.com/billziss-gh/golib/keyring/pkg-overview">Overview</a>
Package keyring implements functions for accessing and storing passwords
in the system's keyring (Keychain on macOS, Credential Manager on
Windows, Secret Service on Linux).




### <a name="github.com/billziss-gh/golib/keyring/pkg-index">Index</a>
* [Constants](#github.com/billziss-gh/golib/keyring/pkg-constants)
* [func Delete(service, user string) error](#github.com/billziss-gh/golib/keyring/Delete)
* [func Get(service, user string) (string, error)](#github.com/billziss-gh/golib/keyring/Get)
* [func Set(service, user, pass string) error](#github.com/billziss-gh/golib/keyring/Set)
* [type FileKeyring](#github.com/billziss-gh/golib/keyring/FileKeyring)
  * [func (self *FileKeyring) Delete(service, user string) error](#github.com/billziss-gh/golib/keyring/FileKeyring.Delete)
  * [func (self *FileKeyring) Get(service, user string) (string, error)](#github.com/billziss-gh/golib/keyring/FileKeyring.Get)
  * [func (self *FileKeyring) Set(service, user, pass string) error](#github.com/billziss-gh/golib/keyring/FileKeyring.Set)
* [type Keyring](#github.com/billziss-gh/golib/keyring/Keyring)
* [type SystemKeyring](#github.com/billziss-gh/golib/keyring/SystemKeyring)
  * [func (self *SystemKeyring) Delete(service, user string) error](#github.com/billziss-gh/golib/keyring/SystemKeyring.Delete)
  * [func (self *SystemKeyring) Get(service, user string) (string, error)](#github.com/billziss-gh/golib/keyring/SystemKeyring.Get)
  * [func (self *SystemKeyring) Set(service, user, pass string) error](#github.com/billziss-gh/golib/keyring/SystemKeyring.Set)


##### <a name="github.com/billziss-gh/golib/keyring/pkg-files">Package files</a>
[keyring_darwin.go](keyring/keyring_darwin.go) [keyring_default.go](keyring/keyring_default.go) [keyring_file.go](keyring/keyring_file.go) 


### <a name="github.com/billziss-gh/golib/keyring/pkg-constants">Constants</a>
``` go
const ErrKeyring = "ErrKeyring"
```



### <a name="github.com/billziss-gh/golib/keyring/Delete">func</a> [Delete](keyring/keyring_default.go#L47)
``` go
func Delete(service, user string) error
```
Delete deletes the password for a service and user in the default keyring.



### <a name="github.com/billziss-gh/golib/keyring/Get">func</a> [Get](keyring/keyring_default.go#L37)
``` go
func Get(service, user string) (string, error)
```
Get gets the password for a service and user in the default keyring.



### <a name="github.com/billziss-gh/golib/keyring/Set">func</a> [Set](keyring/keyring_default.go#L42)
``` go
func Set(service, user, pass string) error
```
Set sets the password for a service and user in the default keyring.




### <a name="github.com/billziss-gh/golib/keyring/FileKeyring">type</a> [FileKeyring](keyring/keyring_file.go#L27)
``` go
type FileKeyring struct {
    Path string
    Key  []byte
    // contains filtered or unexported fields
}
```
FileKeyring is a keyring that stores passwords in a file.










#### <a name="github.com/billziss-gh/golib/keyring/FileKeyring.Delete">func</a> (\*FileKeyring) [Delete](keyring/keyring_file.go#L112)
``` go
func (self *FileKeyring) Delete(service, user string) error
```



#### <a name="github.com/billziss-gh/golib/keyring/FileKeyring.Get">func</a> (\*FileKeyring) [Get](keyring/keyring_file.go#L69)
``` go
func (self *FileKeyring) Get(service, user string) (string, error)
```



#### <a name="github.com/billziss-gh/golib/keyring/FileKeyring.Set">func</a> (\*FileKeyring) [Set](keyring/keyring_file.go#L87)
``` go
func (self *FileKeyring) Set(service, user, pass string) error
```



### <a name="github.com/billziss-gh/golib/keyring/Keyring">type</a> [Keyring](keyring/keyring_default.go#L22)
``` go
type Keyring interface {
    // Get gets the password for a service and user.
    Get(service, user string) (string, error)

    // Set sets the password for a service and user.
    Set(service, user, pass string) error

    // Delete deletes the password for a service and user.
    Delete(service, user string) error
}
```
Keyring is the interface that a system-specific or custom keyring must
implement.


``` go
var DefaultKeyring Keyring
```
The default keyring.










### <a name="github.com/billziss-gh/golib/keyring/SystemKeyring">type</a> [SystemKeyring](keyring/keyring_darwin.go#L24)
``` go
type SystemKeyring struct {
}
```
SystemKeyring implements the system-specific keyring.










#### <a name="github.com/billziss-gh/golib/keyring/SystemKeyring.Delete">func</a> (\*SystemKeyring) [Delete](keyring/keyring_darwin.go#L58)
``` go
func (self *SystemKeyring) Delete(service, user string) error
```



#### <a name="github.com/billziss-gh/golib/keyring/SystemKeyring.Get">func</a> (\*SystemKeyring) [Get](keyring/keyring_darwin.go#L27)
``` go
func (self *SystemKeyring) Get(service, user string) (string, error)
```



#### <a name="github.com/billziss-gh/golib/keyring/SystemKeyring.Set">func</a> (\*SystemKeyring) [Set](keyring/keyring_darwin.go#L49)
``` go
func (self *SystemKeyring) Set(service, user, pass string) error
```










----
## <a name="github.com/billziss-gh/golib/retry">Package retry</a>
_[[godoc.org](https://godoc.org/github.com/billziss-gh/golib/retry)]_

`import "github.com/billziss-gh/golib/retry"`

* [Overview](#github.com/billziss-gh/golib/retry/pkg-overview)
* [Index](#github.com/billziss-gh/golib/retry/pkg-index)

### <a name="github.com/billziss-gh/golib/retry/pkg-overview">Overview</a>
Package retry implements simple retry functionality.

For example to retry an HTTP request:


	func Do(client *http.Client, req *http.Request) (rsp *http.Response, err error) {
	    retry.Retry(
	        retry.Count(5),
	        retry.Backoff(time.Second, time.Second*30),
	        func(i int) bool {
	            if 0 < i {
	                req.Body, err = req.GetBody()
	                if nil != err {
	                    return false
	                }
	            }
	            rsp, err = client.Do(req)
	            if nil != err {
	                return false
	            }
	            if 500 <= rsp.StatusCode && nil != req.GetBody {
	                rsp.Body.Close()
	                return true
	            }
	            return false
	        })
	
	    return
	}




### <a name="github.com/billziss-gh/golib/retry/pkg-index">Index</a>
* [func Backoff(sleep, maxsleep time.Duration) func(int) bool](#github.com/billziss-gh/golib/retry/Backoff)
* [func Count(retries int) func(int) bool](#github.com/billziss-gh/golib/retry/Count)
* [func Retry(actions ...func(int) bool)](#github.com/billziss-gh/golib/retry/Retry)


##### <a name="github.com/billziss-gh/golib/retry/pkg-files">Package files</a>
[retry.go](retry/retry.go) 





### <a name="github.com/billziss-gh/golib/retry/Backoff">func</a> [Backoff](retry/retry.go#L67)
``` go
func Backoff(sleep, maxsleep time.Duration) func(int) bool
```
Backoff implements an exponential backoff with jitter.



### <a name="github.com/billziss-gh/golib/retry/Count">func</a> [Count](retry/retry.go#L60)
``` go
func Count(retries int) func(int) bool
```
Count limits the number of retries performed by Retry.



### <a name="github.com/billziss-gh/golib/retry/Retry">func</a> [Retry](retry/retry.go#L49)
``` go
func Retry(actions ...func(int) bool)
```
Retry performs actions repeatedly until one of the actions returns false.











----
## <a name="github.com/billziss-gh/golib/trace">Package trace</a>
_[[godoc.org](https://godoc.org/github.com/billziss-gh/golib/trace)]_

`import "github.com/billziss-gh/golib/trace"`

* [Overview](#github.com/billziss-gh/golib/trace/pkg-overview)
* [Index](#github.com/billziss-gh/golib/trace/pkg-index)

### <a name="github.com/billziss-gh/golib/trace/pkg-overview">Overview</a>
Package trace provides a simple tracing facility for Go functions.
Given the function below, program execution will be traced whenever
the function is entered or exited.


	func fn(p1 ptype1, p2 ptype2, ...) (r1 rtyp1, r2 rtype2, ...) {
	    defer trace.Trace(0, "TRACE", p1, p2)(&r1, &r2)
	    // ...
	}

The trace facility is disabled unless the variable Verbose is true and
the environment variable GOLIB_TRACE is set to a pattern matching one
of the traced functions. A pattern is a a comma-separated list of
file-style patterns containing wildcards such as * and ?.




### <a name="github.com/billziss-gh/golib/trace/pkg-index">Index</a>
* [Variables](#github.com/billziss-gh/golib/trace/pkg-variables)
* [func Trace(skip int, prfx string, vals ...interface{}) func(vals ...interface{})](#github.com/billziss-gh/golib/trace/Trace)
* [func Tracef(skip int, form string, vals ...interface{})](#github.com/billziss-gh/golib/trace/Tracef)


##### <a name="github.com/billziss-gh/golib/trace/pkg-files">Package files</a>
[trace.go](trace/trace.go) 



### <a name="github.com/billziss-gh/golib/trace/pkg-variables">Variables</a>
``` go
var (
    Verbose = false
    Pattern = os.Getenv("GOLIB_TRACE")
)
```


### <a name="github.com/billziss-gh/golib/trace/Trace">func</a> [Trace](trace/trace.go#L131)
``` go
func Trace(skip int, prfx string, vals ...interface{}) func(vals ...interface{})
```


### <a name="github.com/billziss-gh/golib/trace/Tracef">func</a> [Tracef](trace/trace.go#L162)
``` go
func Tracef(skip int, form string, vals ...interface{})
```










----
## <a name="github.com/billziss-gh/golib/util">Package util</a>
_[[godoc.org](https://godoc.org/github.com/billziss-gh/golib/util)]_

`import "github.com/billziss-gh/golib/util"`

* [Overview](#github.com/billziss-gh/golib/util/pkg-overview)
* [Index](#github.com/billziss-gh/golib/util/pkg-index)

### <a name="github.com/billziss-gh/golib/util/pkg-overview">Overview</a>
Package util contains general utility functions.




### <a name="github.com/billziss-gh/golib/util/pkg-index">Index</a>
* [func ReadAeData(path string, key []byte) (data []byte, err error)](#github.com/billziss-gh/golib/util/ReadAeData)
* [func ReadData(path string) (data []byte, err error)](#github.com/billziss-gh/golib/util/ReadData)
* [func ReadFunc(path string, fn func(*os.File) (interface{}, error)) (data interface{}, err error)](#github.com/billziss-gh/golib/util/ReadFunc)
* [func WriteAeData(path string, perm os.FileMode, data []byte, key []byte) (err error)](#github.com/billziss-gh/golib/util/WriteAeData)
* [func WriteData(path string, perm os.FileMode, data []byte) (err error)](#github.com/billziss-gh/golib/util/WriteData)
* [func WriteFunc(path string, perm os.FileMode, fn func(*os.File) error) (err error)](#github.com/billziss-gh/golib/util/WriteFunc)


##### <a name="github.com/billziss-gh/golib/util/pkg-files">Package files</a>
[doc.go](util/doc.go) [ioae.go](util/ioae.go) [ioutil.go](util/ioutil.go) 





### <a name="github.com/billziss-gh/golib/util/ReadAeData">func</a> [ReadAeData](util/ioae.go#L24)
``` go
func ReadAeData(path string, key []byte) (data []byte, err error)
```


### <a name="github.com/billziss-gh/golib/util/ReadData">func</a> [ReadData](util/ioutil.go#L36)
``` go
func ReadData(path string) (data []byte, err error)
```


### <a name="github.com/billziss-gh/golib/util/ReadFunc">func</a> [ReadFunc](util/ioutil.go#L23)
``` go
func ReadFunc(path string, fn func(*os.File) (interface{}, error)) (data interface{}, err error)
```


### <a name="github.com/billziss-gh/golib/util/WriteAeData">func</a> [WriteAeData](util/ioae.go#L56)
``` go
func WriteAeData(path string, perm os.FileMode, data []byte, key []byte) (err error)
```


### <a name="github.com/billziss-gh/golib/util/WriteData">func</a> [WriteData](util/ioutil.go#L74)
``` go
func WriteData(path string, perm os.FileMode, data []byte) (err error)
```


### <a name="github.com/billziss-gh/golib/util/WriteFunc">func</a> [WriteFunc](util/ioutil.go#L48)
``` go
func WriteFunc(path string, perm os.FileMode, fn func(*os.File) error) (err error)
```







