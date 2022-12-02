# @brief 获取${root_dir}下所有的目录名（非递归）
FUNCTION(list_dirs root_dir output_dirs)
    IF((NOT EXIST ${root_dir}) OR (NOT IS_DIRECTORY ${root_dir}))
        MESSAGE(FATAL_ERROR "root_dir: ${root_dir} invalid")
    ENDIF()

    SET(TMP_OUT )

    FILE(GLOB TMP_PATHS ${root_dir}/*)
    FOREACH(var IN LISTS TMP_PATHS)
        IF(IS_DIRECTORY ${var})
            LIST(APPEND TMP_OUT ${var})
        ENDIF()
    ENDFOREACH()

    SET(${output_dirs} ${TMP_OUT} PARENT_SCOPE)
ENDFUNCTION()

# @brief 获取${root_dir}下所有的子目录（递归）
FUNCTION(list_dirs_recursive root_dir output_dirs)
    SET(TMP_OUT ${${output_dirs}})

    list_dirs(${root_dir} TMP_SUBS)

    FOREACH(var IN LISTS TMP_SUBS)
        list_dirs_recursive(${var} TMP_OUT "recursive")
    ENDFOREACH()

    IF("${ARGN}" STREQUAL "recursive")
        SET(${output_dirs} ${TMP_OUT} ${root_dir} PARENT_SCOPE)
    ELSE()
        SET(${output_dirs} ${TMP_OUT} PARENT_SCOPE)
    ENDIF()
ENDFUNCTION()

# @brief 获取文件夹下所有的文件（非递归）
FUNCTION(list_files root_dir output_files)
    IF((NOT EXIST ${root_dir}) OR (NOT IS_DIRECTORY ${root_dir}))
        MESSAGE(FATAL_ERROR "root_dir: ${root_dir} invalid")
    ENDIF()

    AUX_SOURCE_DIRECTORY(${root_dir} TMP_OUT_FILES)
    SET(${output_files} ${TMP_OUT_FILES} PARENT_SCOPE)
ENDFUNCTION()

# @brief 获取文件夹下所有的文件（递归）
FUNCTION(list_files_recursive root_dir output_files)
    list_dirs_recursive(${root_dir} list_files_DIRS "recursive")

    SET(TMP_OUT )
    FOREACH(var IN LISTS list_files_DIRS)
        AUX_SOURCE_DIRECTORY(${var} TMP_CUR_FILES)
        LIST(APPEND TMP_OUT ${TMP_CUR_FILES})
    ENDFOREACH()

    SET(${output_files} ${TMP_OUT} PARENT_SCOPE)
ENDFUNCTION()

# @brief 添加头文件目录（递归）
FUNCTION(include_dirs_recursive root_dir)
    list_dirs_recursive(${root_dir} INCLUDE_PATHS "recursive")

    INCLUDE_DIRECTORIES(${INCLUDE_PATHS})
    IF(${DEBUG})
        FOREACH(var IN LISTS INCLUDE_PATHS)
            MESSAGE(STATUS "include path: " ${var})
        ENDFOREACH()
    ENDIF()
ENDFUNCTION()