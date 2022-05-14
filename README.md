# grep-go
a grep command written in go

### Usage:
grep-go can read  from a file(dir) or a pipe.

### Help:
```
NAME:
   grep - a grep written in go. just for study linux and go.

USAGE:
   grep [global options] command [command options] grep PATTERNS [FILE...]
       grep -e  PATTERNS ... [FILE...]
       grep -f -r PATTERN_FILE ... [FILE...]


VERSION:
   0.0.1

AUTHOR:
   tanqiangyes <826285820@qq.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --regexp, -e       This flag treats patterns as regular expressions and looks for content in the corresponding file that matches the regular expression. (default: false)
   --recursive, -r    Whether to look recursively in the path. (default: false)
   --file, -f         get pattern from the file. (default: false)
   --line-number, -n  print the number of lines matched. (default: true)
   --ignore-case, -i  Ignore  case  distinctions,  so that characters that differ, only in case match each other. (default: false)
   --help, -h         show help (default: false)
   --version, -v      print the version (default: false)
```

### Future:
- color for every suit text
- beautification code 
- future add
- add tests for code