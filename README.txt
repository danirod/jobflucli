jobflucli
=========

    A Go application for browsing the JobFluent public feed through a
    terminal using an UI that ressembles the UI of other terminal
    applications such as Mutt.

    Looking desperately for a job to pay your bills, but you don't want
    to give up on your fancy desires to use a terminal at all times?
    Want to sneak out what's going on around you but you don't want to
    open suspicious tabs on your work computer?  This app is for you,
    my friend.

Installation
============

    Make sure you have a recent Go distribution installed on your system.

    Then run in your terminal

        go get -u danirod.es/pkg/jobflucli

    To retrieve the source code and compile the application. Provided that
    your PATH is set properly, you will then be able to run the application
    by executing

        $ jobflucli

    On your UNIX or GNU/Linux terminal, or

        C:\> jobflucli

    On your Windows cmd.exe or powershell.exe system prompt.

Source code
===========

    Visit https://git.danirod.es/jobflucli.git

    There is an additional mirror, for all those folks who insist on
    thinking that GitHub == Git at https://github.com/danirod/jobflucli.
    I don't like centralizing stuff on GitHub, but if you want to feed me
    with GitHub stars and karma, I won't hate it.

Disclaimer
==========

    jobflucli fetches public data provided by JobFluent HTTP Feeds.

    This application is totally not related to JobFluent. The fact that they
    even bother to provide JSON and XML feeds prooves that they are good
    Internet citizens. Make sure to visit their website sometimes, however.

    Please, don't use this application in places where you are not allowed.
    Don't point at me if something bad happens when you use this application.

    I made this application in less than 48 hours and it's my first Go
    project. Please, forgive any errors this application may have. I hope to
    have them fixed.
