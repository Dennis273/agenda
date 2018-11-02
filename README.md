# Agenda

## Overview

Agenda is an CLI-based meeting-managing application which supports multiple users.

## Installing

## Getting Started

### Users

#### register

#### login

#### logout

#### deleteUser

#### queryUser

### Meetings

#### createMeeting

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

##### 创建一个会议(需要先登陆)

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

##### 不能创建同名会议

```bash
$ agenda createMeeting -t meeting1 -s 2018-11-02-20-00 -e
...
Create meeting failed: meeting exit
```

##### 参与者不能有相同的用户, 发起者不能作为参与者

```bash
$ agenda createMeeting -t meeting2  -s 2018-11-02-20-10 -e 2018-11-02-20-20 -p user2 -p user2
...
Create meeting failed: You could not repeatly attend a user to a same meeting

$ agenda createMeeting -t meeting2  -s 2018-11-02-20-10 -e 2018-11-02-20-20 -p user1
...
Create meeting failed: You could not attend this meeting as participator
```

##### 参与者必须为已注册用户

```bash
$ agenda createMeeting -t meeting2  -s 2018-11-02-20-10 -e 2018-11-02-20-20 -p errorUser
...
Create meeting failed: errorUser is not exit
```

##### 用户发起者或者参与者不能分身参与会议(允许时间端点重叠)

```bash
$ agenda createMeeting -t meeting2  -s 2018-11-02-20-09 -e 2018-11-02-20-20 -p user2
...
Create meeting failed: You are busy at that time
```

登陆user4

```bash
$ agenda createMeeting -t meeting2  -s 2018-11-02-20-09 -e 2018-11-02-20-20 -p user2
...
Create meeting failed: user2 is busy at that time
```

#### modifyMember

```bash
Usage:
  Agenda modifyMember [flags]

Flags:
  -a, --add stringArray      add member(s) to the meeting
  -d, --delete stringArray   delete member(s) from the meeting
  -h, --help                 help for modifyMember
  -t, --title string         meeting title
```

##### 向已存在的添加成员或移出成员

```bash
$ agenda modifyMember -t meeting1 -a user4
...
old participators:ModifyMeeting success
```

##### 操作的会议须已创建

```bash
$ agenda modifyMember -t errorMeeting -a user4
...
ModifyMeeting failed: Meeting errorMeeting is not exit
```

##### 添加或移除的成员需要是已注册用户

```bash
$ agenda modifyMember -t meeting1 -a erroruser
...
ModifyMeeting failed: User erroruser is not exit
```

##### 你不能将自己添加到会议中或从会议中移除, 不能一次重复添加或移除成员

```bash
$ agenda modifyMember -t meeting1 -a user1
modifyMember called
meeting title:  meeting1
new participators: 0:user1
old participators:ModifyMeeting failed: You are the holder of this meeting

$ agenda modifyMember -t meeting1 -a user4 -a user4
...
ModifyMeeting failed: user4 was a participator of this meeting
```

##### 添加的成员之前不在在会议中, 移除的成员之前要在会议中

```bash
$ agenda modifyMember -t meeting1 -a user2
...
ModifyMeeting failed: user2 was a participator of this meeting

$ agenda modifyMember -t meeting1 -d user5
...
ModifyMeeting failed: user5 was not a participator of this meeting
```

##### 你必须为创建者

登陆user2

```bash
$ agenda modifyMember -t meeting1 -a user4
...
ModifyMeeting failed: You are not the holder of meeting meeting1
```

##### 如果会议中没有参与者, 则会议会被取消

```bash
$ agenda modifyMember -t meeting1 -d user2 -d user3 -d user4
...
ModifyMeeting success

$ agenda modifyMember -t meeting1 -a user2
...
ModifyMeeting failed: Meeting meeting1 is not exit
```

##### 向会议添加的成员需要有时间参与会议

```bash
$ agenda modifyMember -t meeting1 -a user3
...
ModifyMeeting failed: user3 is busy at that time
```

#### queryMeeting

```bash
Usage:
  Agenda queryMeeting [flags]

Flags:
  -e, --endTime string     the end time(yyyy-MM-dd-hh-mm)
  -h, --help               help for queryMeeting
  -s, --startTime string   the start time(yyyy-MM-dd-hh-mm)
```

##### 时间区间中开始时间需要小于等于结束时间

```bash
$ agenda queryMeeting -s 2018-11-02-20-00 -e 2018-11-02-20-01
queryMeeting called
the start time: 2018-11-02-20-00
the end time: 2018-11-02-20-01
StartTime       EndTime         Title           Holder  Participators
2018-11-2-20-0  2018-11-2-20-10 meeting1        user1   user2
Query meeting success
```

#### cancelMeeting

#### quitMeeting

#### clearMeeting