# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.18

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Disable VCS-based implicit rules.
% : %,v


# Disable VCS-based implicit rules.
% : RCS/%


# Disable VCS-based implicit rules.
% : RCS/%,v


# Disable VCS-based implicit rules.
% : SCCS/s.%


# Disable VCS-based implicit rules.
% : s.%


.SUFFIXES: .hpux_make_needs_suffix_list


# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /snap/cmake/619/bin/cmake

# The command to remove a file.
RM = /snap/cmake/619/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/colin/Workspace/src/sc/src

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/colin/Workspace/src/sc/build/sc

# Include any dependencies generated for this target.
include CMakeFiles/sc.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/sc.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/sc.dir/flags.make

CMakeFiles/sc.dir/sc.obj: CMakeFiles/sc.dir/flags.make
CMakeFiles/sc.dir/sc.obj: /home/colin/Workspace/src/sc/src/sc.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/colin/Workspace/src/sc/build/sc/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/sc.dir/sc.obj"
	/usr/local/platone.cdt/bin/platone-cpp $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/sc.dir/sc.obj -c /home/colin/Workspace/src/sc/src/sc.cpp

CMakeFiles/sc.dir/sc.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/sc.dir/sc.i"
	$(CMAKE_COMMAND) -E cmake_unimplemented_variable CMAKE_CXX_CREATE_PREPROCESSED_SOURCE

CMakeFiles/sc.dir/sc.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/sc.dir/sc.s"
	$(CMAKE_COMMAND) -E cmake_unimplemented_variable CMAKE_CXX_CREATE_ASSEMBLY_SOURCE

# Object files for target sc
sc_OBJECTS = \
"CMakeFiles/sc.dir/sc.obj"

# External object files for target sc
sc_EXTERNAL_OBJECTS =

sc.wasm: CMakeFiles/sc.dir/sc.obj
sc.wasm: CMakeFiles/sc.dir/build.make
sc.wasm: CMakeFiles/sc.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/colin/Workspace/src/sc/build/sc/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking CXX executable sc.wasm"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/sc.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/sc.dir/build: sc.wasm

.PHONY : CMakeFiles/sc.dir/build

CMakeFiles/sc.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/sc.dir/cmake_clean.cmake
.PHONY : CMakeFiles/sc.dir/clean

CMakeFiles/sc.dir/depend:
	cd /home/colin/Workspace/src/sc/build/sc && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/colin/Workspace/src/sc/src /home/colin/Workspace/src/sc/src /home/colin/Workspace/src/sc/build/sc /home/colin/Workspace/src/sc/build/sc /home/colin/Workspace/src/sc/build/sc/CMakeFiles/sc.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/sc.dir/depend

