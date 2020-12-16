#!/usr/bin/env node

// dpw@alameda.local
// 2015.03.04
'use strict';

const fs = require('fs'),
    spawn = require('child_process').spawn,
    clearScreen = '[H[2J';

let files = new Set();
let tid = null;

const run = function() {
    process.stdout.write( clearScreen ); 
    console.log('Changed files: ', files);

    let list = Array.from(files);

    const loop = function() {
        const filename = list.shift();
        if (filename) {
            console.log(filename);
            let runner = spawn( 'go', [ 'run', filename ] );

            runner.stdout.on('data', function( data ) {
                process.stdout.write( data );
            });

            runner.stderr.on('data', function( data ) {
                process.stdout.write( data );
            });

            runner.on('close', function(code) {
                loop();
            });
        } else {
            tid = null;
            files.clear();
            console.log('\ntests complete...')
        }
    }

    loop();
};

const changeHandler = function(event, filename) {
    // console.log( 'raw file change: ', filename);
    if ( filename.endsWith('.go') ) {
        files.add( filename );

        if (!tid) {
            tid = setTimeout(function() {
                run();
            }, 250);
        }
    }
};

// run();
fs.watch( './', { recursive:false }, changeHandler );

