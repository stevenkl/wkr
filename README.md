# Wkr

> Nach kurzer Zeit muss das Projekt bereits überarbeitet werden, da die Zielsetzung nicht klar ist/war.


Wkr is simple job-runner, inspired by [Workr](https://github.com/sirikon/workr) but written in Go.
It is primary run on Windows, but running on Linux/MacOS is also planed.


## Purpose

Wkr soll die Möglichkeit bieten selbst definierte Task/Aufgaben starten und Überwachen zu können.
Die Aufgaben werden in einer [Konfigurationsdatei](#configfile) beschrieben,
Wkr zeigt eine Auflistung der Aufgaben an und bietet die Möglichkeit diese per Knopfdruck zu starten.


## Konfigurationsdatei <a link="configfile"></a>

Die Konfigurationsdatei in einem TCL-gültigen Format geschrieben, zum parsen kommt [stevenkl/tcl.go](https://github.com/stevenkl/tcl.go) zum Einsatz.
In der Datei steht nur eine begrenzte Auswahl an Commands zur Verfügung:
* `server` - required, only once
	* `host` - required, only once
	* `port` - required, only once
* `storage` -required, only once
	* `path` - required, only once
* `user` - optional, multiple times
	* `name` - required
	* `password [-raw]` -required, if option `-raw` is set Wkr generates a bcyrpt hash from the password. Otherwise Wkr asumes that the password is already hashed
	* `group` - optional, defaults to "guest"
* `job` - optional, multiple times
	* `name` - required
	* `workdir` -optional, defaults to current directory + jobname
	* `run` - required, the command to run when an execution of the job is triggered


In each of the main sections (server, storage, user, job) you can use some commands to make it easier for dynamic configuration:

* `env <VAR_NAME>` - Returns the value of the given environment variable, or an empty string

### Parse & Validate

The configfile must at first be parsed, setting all values from it to the state config.
After that, some validation is needed. The validation is done by 4 functions:

* `validateServerConfig`
* `validateStorageConfig`
* `validateUsersConfig`
* `validateJobsConfig`

This validation is needed to make sure that `Wkr` can properbly execute with the given config state.
As an example: It should exist min. one user, otherwise Wkr wouldn't be useful.
The same goes for the jobs, if there is no job defined, there is no reason why Wkr should run.


## Build

To build you need [Tsk](https://github.com/stevenkl/tsk) installed:

```shell
go get -u github.com/stevenkl/tsk/cmd/tsk
```

After that call `tsk build`, you can find the executable at `./build/wkr(.exe)`.


## Ablaufplan

### Start ohne Parameter

* Wkr.exe wird ausgeführt
* Laden der Konfigurationsdatei (Default: ./wkr.config)
	* Mit Parameter `-c,--config` kann der Pfad zur Konfigurationsdatei angegeben werden
* Verzeichnisse für die einzelnen Jobs anlegen
* Benutzer prüfen, mindestens ein Benutzer muss in der Gruppe `admin` eingetragen sein
* API-Server starten


### Start mit Parameter `--generate-hash`
* `wkr.exe --generate-hash <password>` wird ausgeführt
* Bcrypt hash für `password` erstellen
* Auf Konsole ausgeben


### Start mit Parameter `--validate-hash`
* `wkr.exe --validate-hash <password> <hash>`
* Bcrypt Passwort mit Hash vergleichen
* Ergebnis auf Konsole ausgeben



## Internal

Zu Beginn wird der State initialisiert, das Laden der Konfiguration füllt diesen State.
Ziel ist es bei jedem Start mit der selben Konfiguration auch das gleiche Verhalten zu bewirken.


### API-Routes

* `GET  /ping`
	* `response` `204 No Content`
* `POST /login`
	* `request` `{"name":"admin","password":"my-password"}`
	* `response` `200 OK {"status":"success","data":"<the-token>"}`
	* `response` `401 Unauthorized {"status":"error":"data":"Unknown credentials"}`
* `GET  /jobs/`
* `GET  /jobs/:job_id`
* `POST /jobs/:job_id/run`
* `GET  /jobs/:job_id/:run_id`
* `GET  /jobs/:job_id/:run_id/stream`

