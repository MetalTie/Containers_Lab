# Containers_Lab
Modern languages provide convenient container libraries (lists, maps, sets, etc.). For this lab, you may not use library containers on the core data-structure paths. In Go, that means on the main container logic you must not rely on maps, container/list, or sorting helpers to maintain order. Instead, implement your own circular list and ordering logic.

Goal: Develop the technology to implement container behavior using only basic constructs. You will implement and compare:

    C: static allocation with two arrays and per-JobType circular lists;
    B: dynamic allocation for a homogeneous record (cars) with per-JobType rings;
    A: dynamic allocation with interfaces for vehicles (car or plane), employees grouped by JobType;
    A+: extend A with multiple vehicles per employee and one extra vehicle type.

Submission (all options):

    Cover sheet first (name, course + section + meeting days, submission date, grading option completed).
    Any required screenshots next, followed by program results, and then the complete code.
    All input/output should be plain text (ASCII). Use traditional file I/O or I/O redirection.

Due:
    Friday, October 30 2026