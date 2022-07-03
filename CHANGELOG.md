# CHANGELOG

# 03/07/2022

## Feature

Include Godocs to gcode package.

Include simple unit test to gcode package.

## Plans

I going to begin to use *Github Projects* to manage issues and plans. Hence, this project will have various changes.
In principle, I will probe the new workflow proposed by *Github* using issues and pull requests with the support of *Github Projects* and his tables, graphs, etc...

Then, I will review how I go to implement the address package. It requires a critical decision that defines the paradigm of the project.

Later, I will complete the current packages with his respective test and godocs, before of realise serious changes to the repositories and the project.

The target is to make a CLI tool that allows the management of any 3d printer based gcode file. For example, a simple subcommand that applies transforms on coordinates like a correction skew.

# 25/06/2022

## Feature

I reimplement package structs and, I create new structs to manage each of the elements that compose a gcode' line.

A block will be composed of multiple gcode' struct that represent each part of the command line, with exception of the comment section.

I do not implement a system to validate a block using some map of allowed commands yet.

Check struct consist in a interface that is implement for diferents algorithms of verification like checksum or crc.

In fact, crc algorithm is not implemented ready, solo exists as a dummy sketch.

Address package allows will construction of an address struct from multiple data types.

Same time, block struct expose different format of representation that include or ommite comments, check or data.

## Plans

I will complete each package with the respective documentation using gdoc system, later.
