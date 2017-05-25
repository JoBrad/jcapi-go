## JCAPI-Go
### Prerequisites ###
#### Java 1.8 ####
As of swagger-codegen 2.2.2, you must use JDK 1.8 or higher. On Debian/Ubuntu:
```
$ sudo add-apt-repository ppa:webupd8team/java
$ sudo apt update; sudo apt-get install oracle-java8-installer
```
The easiest way to get java/swagger-codegen on MacOSX is below.  It
will also setup your `JAVA_HOME` and `PATH` for you.
```
brew cask install java
brew install swagger-codegen
```
If you had to install Java the hard way on your Mac, you should just have to add the appropriate `JAVA_HOME` to your path:
```
$ export JAVA_HOME=`/usr/libexec/java_home -v 1.8`
$ export PATH=${JAVA_HOME}/bin:$PATH
```
Check your version:
```
$ java -version
java version "1.8.0_121"
Java(TM) SE Runtime Environment (build 1.8.0_121-b13)
Java HotSpot(TM) 64-Bit Server VM (build 25.121-b13, mixed mode)
```
#### Swagger-codegen ####
A pre-built version of swagger-codegen is included in this repo.
For more information, updates, or alternate installation methods, please see the github repo: https://github.com/swagger-api/swagger-codegen

Check swagger-codegen:
```
$ java -jar swagger-codegen-cli.jar

OR

$ swagger-codegen
```
This should get you something like the following output.
```
$ Available languages: [android, aspnet5, aspnetcore, async-scala, .... ]
```

#### Generating the API Client

To generate the APIs, run the command below.  It assumes that you have followed the standard practice of laying out the dev env under `$HOME/workspace` and have cloned SI into the "workspace" root.  Alternately you can specify the SI path on the cmdline.

```
$ make all

OR

$ make all SWAGGER_FILE_PATH=/Users/yourmom/mySIdirectory/routes/webui/api
```

If you are developing the APIs, you can clean up by running `make clean`

#### Usage Examples

```
import (
	"github.com/TheJumpCloud/jcapi-go/jcapiv1"
  "github.com/TheJumpCloud/jcapi-go/jcapiv2"
)
...
apiV1 = jcapiv1.NewDefaultApi()
groupsApiV2 = jcapiv2.NewUserGroupsApi()
```
