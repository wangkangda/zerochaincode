# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 2.8

#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:

# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list

# Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:
.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/wkdisee/golang/work/src/test/gocoin/libzerocoin

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/wkdisee/golang/work/src/test/gocoin/lib

# Include any dependencies generated for this target.
include CMakeFiles/paramgen.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/paramgen.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/paramgen.dir/flags.make

CMakeFiles/paramgen.dir/paramgen.cpp.o: CMakeFiles/paramgen.dir/flags.make
CMakeFiles/paramgen.dir/paramgen.cpp.o: /home/wkdisee/golang/work/src/test/gocoin/libzerocoin/paramgen.cpp
	$(CMAKE_COMMAND) -E cmake_progress_report /home/wkdisee/golang/work/src/test/gocoin/lib/CMakeFiles $(CMAKE_PROGRESS_1)
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Building CXX object CMakeFiles/paramgen.dir/paramgen.cpp.o"
	/usr/bin/c++   $(CXX_DEFINES) $(CXX_FLAGS) -o CMakeFiles/paramgen.dir/paramgen.cpp.o -c /home/wkdisee/golang/work/src/test/gocoin/libzerocoin/paramgen.cpp

CMakeFiles/paramgen.dir/paramgen.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/paramgen.dir/paramgen.cpp.i"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_FLAGS) -E /home/wkdisee/golang/work/src/test/gocoin/libzerocoin/paramgen.cpp > CMakeFiles/paramgen.dir/paramgen.cpp.i

CMakeFiles/paramgen.dir/paramgen.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/paramgen.dir/paramgen.cpp.s"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_FLAGS) -S /home/wkdisee/golang/work/src/test/gocoin/libzerocoin/paramgen.cpp -o CMakeFiles/paramgen.dir/paramgen.cpp.s

CMakeFiles/paramgen.dir/paramgen.cpp.o.requires:
.PHONY : CMakeFiles/paramgen.dir/paramgen.cpp.o.requires

CMakeFiles/paramgen.dir/paramgen.cpp.o.provides: CMakeFiles/paramgen.dir/paramgen.cpp.o.requires
	$(MAKE) -f CMakeFiles/paramgen.dir/build.make CMakeFiles/paramgen.dir/paramgen.cpp.o.provides.build
.PHONY : CMakeFiles/paramgen.dir/paramgen.cpp.o.provides

CMakeFiles/paramgen.dir/paramgen.cpp.o.provides.build: CMakeFiles/paramgen.dir/paramgen.cpp.o

# Object files for target paramgen
paramgen_OBJECTS = \
"CMakeFiles/paramgen.dir/paramgen.cpp.o"

# External object files for target paramgen
paramgen_EXTERNAL_OBJECTS =

paramgen: CMakeFiles/paramgen.dir/paramgen.cpp.o
paramgen: CMakeFiles/paramgen.dir/build.make
paramgen: libzerocoin.so
paramgen: /usr/lib/x86_64-linux-gnu/libboost_system.so
paramgen: /usr/lib/x86_64-linux-gnu/libcrypto.so
paramgen: CMakeFiles/paramgen.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --red --bold "Linking CXX executable paramgen"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/paramgen.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/paramgen.dir/build: paramgen
.PHONY : CMakeFiles/paramgen.dir/build

CMakeFiles/paramgen.dir/requires: CMakeFiles/paramgen.dir/paramgen.cpp.o.requires
.PHONY : CMakeFiles/paramgen.dir/requires

CMakeFiles/paramgen.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/paramgen.dir/cmake_clean.cmake
.PHONY : CMakeFiles/paramgen.dir/clean

CMakeFiles/paramgen.dir/depend:
	cd /home/wkdisee/golang/work/src/test/gocoin/lib && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/wkdisee/golang/work/src/test/gocoin/libzerocoin /home/wkdisee/golang/work/src/test/gocoin/libzerocoin /home/wkdisee/golang/work/src/test/gocoin/lib /home/wkdisee/golang/work/src/test/gocoin/lib /home/wkdisee/golang/work/src/test/gocoin/lib/CMakeFiles/paramgen.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/paramgen.dir/depend

