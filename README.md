# jsonfmt

accept loosely formatted json, output strictly formatted json

## Install / Update

```
go get -f -u -v github.com/phinze/jsonfmt
```

## Usage

```
jsonfmt <json_file>
```

Reformats the file in place.

## Using with VIM

In your `.vimrc`:

```vim
function! JsonFormat()
  " save cursor position and many other things
  let l:curw=winsaveview()

  " Write current unsaved buffer to a temp file
  let l:tmpname = tempname()
  call writefile(getline(1, '$'), l:tmpname)

  " populate the final command with user based fmt options
  let command = "jsonfmt"

  " execute our command...
  let out = system(command . " " . l:tmpname)

  if v:shell_error == 0
    " remove undo point caused via BufWritePre
    try | silent undojoin | catch | endtry

    " Replace current file with temp file, then reload buffer
    call rename(l:tmpname, expand('%'))
    silent edit!
    let &syntax = &syntax
  endif

  " restore our cursor/windows positions
  call winrestview(l:curw)
endfunction

autocmd BufWritePre *.json call JsonFormat()
```

## Acknowledgements

 * Uses [@daviddengcn](https://github.com/daviddengcn)'s
   [ljson](https://github.com/daviddengcn/ljson) for the loose
   input parsing a.k.a. the actually difficult part ;D
 * Vim code stolen from [@fatih](https://github.com/fatih)'s delightful
   [vim-go](https://github.com/fatih/vim-go) project.
