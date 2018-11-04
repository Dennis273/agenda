# Agenda

## Overview

Agenda is an CLI-based meeting-managing application which supports multiple users.

## Getting Started

### Users

#### ` register`

##### Basic Usage

```bash
Usage:
  Agenda register [flags]

Flags:
  -e, --email string         Email
  -h, --help                 help for register
  -p, --password string      Password
  -n, --phoneNumber string   Phone number
  -u, --username string      Username to login
```

#### `login`

```bash
Usage:
  Agenda login [flags]

Flags:
  -h, --help              help for login
  -p, --password string   Password
  -u, --username string   Username
```

* User cannot login another user while logged in

    ```bash
    $ agenda login -uUser -pPass
    Logout success
    $ agenda login -uUser -pPass
    Already logged in as User
    ```

#### `logout`

```bash
Usage:
  Agenda logout [flags]

Flags:
  -h, --help   help for logout
```

* User cannot login another user while not logged in

  ```bash
  // not logged in
  $ agenda logout
  User not logged in.
  ```


#### `deleteUser`

```bash
Usage:
  Agenda deleteUser [flags]

Flags:
  -h, --help   help for deleteUser
```

* Only logged user can delete its account

  ```bash
  // not logged in
  $ agenda deleteUser
  User not logged in
  ```

#### `queryUser`

```bash
Usage:
  Agenda queryUser [flags]

Flags:
  -h, --help   help for queryUser
```

* Only logged user can query all users

  ```bash
  // not logged in
  $ agenda queryUser
  Permission denied
  ```

### Meetings

#### `createMeeting`

```bash
Usage:
  Agenda createMeeting [flags]

Flags:
  -e, --endTime string              the end time of the meeting(yyyy-MM-dd-hh-mm)
  -h, --help                        help for createMeeting
  -p, --participators stringArray   the participators of the meeting
  -s, --startTime string            the time when the meeting begin(yyyy-MM-dd-hh-mm)
  -t, --title string                meeting title
```

* Creating a Meeting (Login required)

    ```bash
    $ agenda createMeeting -t meeting1 -s 2018-11-02-20-00 -e 2018-11-02-20-10 -p user2 -p user3
    createMeeting called
    meeting title:  meeting1
    start time:  2018-11-02-20-00
    end time:  2018-11-02-20-10
    participators:
    0:user2
    1:user3
    Create meeting success
    ```

* You cannot create a meeting with existing name

    ```bash
    $ agenda createMeeting -t meeting1 -s 2018-11-02-20-00 -e
    ...
    Create meeting failed: meeting exit
    ```

* The sponsor cannot be a participator

    ```bash
    $ agenda createMeeting -t meeting2  -s 2018-11-02-20-10 -e 2018-11-02-20-20 -p user2 -p user2
    ...
    Create meeting failed: You could not repeatly attend a user to a same meeting

    $ agenda createMeeting -t meeting2  -s 2018-11-02-20-10 -e 2018-11-02-20-20 -p user1
    ...
    Create meeting failed: You could not attend this meeting as participator
    ```

* Participator should be an registered user

    ```bash
    $ agenda createMeeting -t meeting2  -s 2018-11-02-20-10 -e 2018-11-02-20-20 -p errorUser
    ...
    Create meeting failed: errorUser is not exit
    ```

* Users are not allowed to participate in meetings at the same time

    ```bash
    $ agenda createMeeting -t meeting2  -s 2018-11-02-20-09 -e 2018-11-02-20-20 -p user2
    ...
    Create meeting failed: You are busy at that time

    // login with User4
    $ agenda createMeeting -t meeting2  -s 2018-11-02-20-09 -e 2018-11-	  02-20-20 -p user2
    ...
    reate meeting failed: user2 is busy at that time
    ```


#### `modifyMember`

```bash 
Usage:
  Agenda modifyMember [flags]
Flags:
  -a, --add stringArray      add member(s) to the meeting
  -d, --delete stringArray   delete member(s) from the meeting
  -h, --help                 help for modifyMember
  -t, --title string         meeting title
```


* Update meeting participator

    ```bash
    $ agenda modifyMember -t meeting1 -a user4
    ...
    old participators:ModifyMeeting success
    ```

* Member should be existing user

    ```bash
    $ agenda modifyMember -t meeting1 -a erroruser
    ...
    ModifyMeeting failed: User erroruser is not exit
    ```


* Only sponsor can update a meeting

    ```bash
    // logi as user2
    $ agenda modifyMember -t meeting1 -a user4
    ...
    ModifyMeeting failed: You are not the holder of meeting meeting1
    ```

* A meeting will be canceled if it has no paritcipator

    ```bash
    $ agenda modifyMember -t meeting1 -d user2 -d user3 -d user4
    ...
    ModifyMeeting success

    $ agenda modifyMember -t meeting1 -a user2
    ...
    ModifyMeeting failed: Meeting meeting1 is not exit
    ```

* Cannot add members who are buzy into a meeting

    ```bash
    $ agenda modifyMember -t meeting1 -a user3
    ...
    ModifyMeeting failed: user3 is busy at that time
    ```

#### `queryMeeting`

    ```bash
    Usage:
      Agenda queryMeeting [flags]
    
    Flags:
      -e, --endTime string     the end time(yyyy-MM-dd-hh-mm)
      -h, --help               help for queryMeeting
      -s, --startTime string   the start time(yyyy-MM-dd-hh-mm)
    ```

* Time check

    ```bash
    $ agenda queryMeeting -s 2018-11-02-20-00 -e 2018-11-02-20-01
    queryMeeting called
    the start time: 2018-11-02-20-00
    the end time: 2018-11-02-20-01
    StartTime       EndTime         Title           Holder  Participators
    2018-11-2-20-0  2018-11-2-20-10 meeting1        user1   user2
    Query meeting success
    ```

#### `cancelMeeting`

#### `quitMeeting`

#### `clearMeeting`