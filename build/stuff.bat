@echo off
rem #%L
rem Codenjoy - it's a dojo-like platform from developers to developers.
rem %%
rem Copyright (C) 2012 - 2022 Codenjoy
rem %%
rem This program is free software: you can redistribute it and/or modify
rem it under the terms of the GNU General Public License as
rem published by the Free Software Foundation, either version 3 of the
rem License, or (at your option) any later version.
rem
rem This program is distributed in the hope that it will be useful,
rem but WITHOUT ANY WARRANTY; without even the implied warranty of
rem MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
rem GNU General Public License for more details.
rem
rem You should have received a copy of the GNU General Public
rem License along with this program.  If not, see
rem <http://www.gnu.org/licenses/gpl-3.0.html>.
rem #L%
@echo on

@echo off

if "%RUN%"=="" set RUN=%CD%\run
if "%STUFF%"=="" set STUFF=%CD%\stuff

call %RUN% :init_colors

:check_run_mode
    if "%*"=="" (       
        call :run_executable 
    ) else (
        call :run_library %*
    )
    goto :eof

:run_executable
    rem run stuff.bat as executable script
    call %RUN% :color ‘%CL_INFO%‘ ‘This is not executable script. Please use 'run.bat' only.‘
    call %RUN% :ask   
    goto :eof

:run_library
    rem run stuff.bat as library
    call %*     
    goto :eof          

:settings
    if "%INSTALL_LOCALLY%"=="true" ( set GOPATH=)

    if "%GOPATH%"==""     ( set NO_GO=true)
    if "%NO_GO%"=="true"  ( set GOPATH=%ROOT%\.golang)
    if "%NO_GO%"=="true"  ( set PATH=%GOPATH%\bin;%PATH%)

    set GO=%GOPATH%\bin\go

    echo Language environment variables
    call %RUN% :color ‘%CL_INFO%‘ ‘PATH=%PATH%‘
    call %RUN% :color ‘%CL_INFO%‘ ‘GOPATH=%GOPATH%‘

    set ARCH_URL=https://golang.org/dl/go1.16.5.windows-amd64.zip
    set ARCH_FOLDER=go
    goto :eof

:install
    call %RUN% :install golang %ARCH_URL% %ARCH_FOLDER%
    goto :eof

:version
    call %RUN% :eval_echo_color ‘%GO% version‘
    goto :eof

:build
    rem do nothing
    goto :eof

:test    
    call %RUN% :eval_echo ‘%GO% test .\...‘
    echo.
    goto :eof

:run
    call %RUN% :eval_echo ‘%GO% run .\main.go %GAME_TO_RUN% %SERVER_URL%‘
    goto :eof