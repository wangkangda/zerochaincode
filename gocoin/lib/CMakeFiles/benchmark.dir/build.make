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
include CMakeFiles/benchmark.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/benchmark.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/benchmark.dir/flags.make

CMakeFiles/benchmark.dir/Benchmark.cpp.o: CMakeFiles/benchmark.dir/flags.make
CMakeFiles/benchmark.dir/Benchmark.cpp.o: /home/wkdisee/golang/work/src/test/gocoin/libzerocoin/Benchmark.cpp
	$(CMAKE_COMMAND) -E cmake_progress_report /home/wkdisee/golang/work/src/test/gocoin/lib/CMakeFiles $(CMAKE_PROGRESS_1)
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Building CXX object CMakeFiles/benchmark.dir/Benchmark.cpp.o"
	/usr/bin/c++   $(CXX_DEFINES) $(CXX_FLAGS) -o CMakeFiles/benchmark.dir/Benchmark.cpp.o -c /home/wkdisee/golang/work/src/test/gocoin/libzerocoin/Benchmark.cpp

CMakeFiles/benchmark.dir/Benchmark.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/benchmark.dir/Benchmark.cpp.i"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_FLAGS) -E /home/wkdisee/golang/work/src/test/gocoin/libzerocoin/Benchmark.cpp > CMakeFiles/benchmark.dir/Benchmark.cpp.i

CMakeFiles/benchmark.dir/Benchmark.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/benchmark.dir/Benchmark.cpp.s"
	/usr/bin/c++  $(CXX_DEFINES) $(CXX_FLAGS) -S /home/wkdisee/golang/work/src/test/gocoin/libzerocoin/Benchmark.cpp -o CMakeFiles/benchmark.dir/Benchmark.cpp.s

CMakeFiles/benchmark.dir/Benchmark.cpp.o.requires:
.PHONY : CMakeFiles/benchmark.dir/Benchmark.cpp.o.requires

CMakeFiles/benchmark.dir/Benchmark.cpp.o.provides: CMakeFiles/benchmark.dir/Benchmark.cpp.o.requires
	$(MAKE) -f CMakeFiles/benchmark.dir/build.make CMakeFiles/benchmark.dir/Benchmark.cpp.o.provides.build
.PHONY : CMakeFiles/benchmark.dir/Benchmark.cpp.o.provides

CMakeFiles/benchmark.dir/Benchmark.cpp.o.provides.build: CMakeFiles/benchmark.dir/Benchmark.cpp.o

# Object files for target benchmark
benchmark_OBJECTS = \
"CMakeFiles/benchmark.dir/Benchmark.cpp.o"

# External object files for target benchmark
benchmark_EXTERNAL_OBJECTS =

benchmark: CMakeFiles/benchmark.dir/Benchmark.cpp.o
benchmark: CMakeFiles/benchmark.dir/build.make
benchmark: libzerocoin.so
benchmark: /usr/lib/x86_64-linux-gnu/libboost_system.so
benchmark: /usr/lib/x86_64-linux-gnu/libcrypto.so
benchmark: CMakeFiles/benchmark.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --red --bold "Linking CXX executable benchmark"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/benchmark.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/benchmark.dir/build: benchmark
.PHONY : CMakeFiles/benchmark.dir/build

CMakeFiles/benchmark.dir/requires: CMakeFiles/benchmark.dir/Benchmark.cpp.o.requires
.PHONY : CMakeFiles/benchmark.dir/requires

CMakeFiles/benchmark.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/benchmark.dir/cmake_clean.cmake
.PHONY : CMakeFiles/benchmark.dir/clean

CMakeFiles/benchmark.dir/depend:
	cd /home/wkdisee/golang/work/src/test/gocoin/lib && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/wkdisee/golang/work/src/test/gocoin/libzerocoin /home/wkdisee/golang/work/src/test/gocoin/libzerocoin /home/wkdisee/golang/work/src/test/gocoin/lib /home/wkdisee/golang/work/src/test/gocoin/lib /home/wkdisee/golang/work/src/test/gocoin/lib/CMakeFiles/benchmark.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/benchmark.dir/depend

