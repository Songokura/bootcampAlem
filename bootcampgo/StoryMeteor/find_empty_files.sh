#!/bin/bash                                                                                                                                                     

script_name=$(basename "$0")

# Check if the number of arguments is zero                                                                                                                      
if [ "$#" -eq 0 ]; then
  echo "usage: $script_name dir_path"
  exit 0
fi

# Get the directory path from the first argument                                                                                                                
dir_path="$1"

# Check if the directory exists and is a directory                                                                                                              
if [ ! -d "$dir_path" ]; then
  if [ -e "$dir_path" ]; then
    echo "error: $dir_path is not a directory"
  else
    echo "error: directory $dir_path not found"
  fi
  exit 0
fi

# Find empty files or files with only spaces or empty lines                                                                                                     
find "$dir_path" -type f | while IFS= read -r file; do
  if [ ! -s "$file" ] || [ -z "$(grep -v '^\s*$' "$file")" ]; then
    echo "$file"
  fi
done
