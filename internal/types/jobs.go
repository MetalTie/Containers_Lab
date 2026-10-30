package types

import "strings"

const JobTypeCount = 8

type JobType int

const (
    Accountant JobType = iota
    Analysist
    Manager
    Manufacturing
    Programmer
    Inventory
    Sales
    SoftwareEnginner
)

func ParseJobType(s string) (JobType, bool) {
    switch strings.ToLower(s) {
    case "accountant":
        return Accountant, true
    case "analysist", "analysis":
        return Analysist, true
    case "manager":
        return Manager, true
    case "manufacturing":
        return Manufacturing, true
    case "programmer":
        return Programmer, true
    case "inventory":
        return Inventory, true
    case "sales":
        return Sales, true
    case "softwareengineer", "software_engineer", "software engineer", "softwareenginner":
        return SoftwareEnginner, true
    default:
        return 0, false
    }
}

// String method so JobType prints its text name automatically
func (j JobType) String() string {
    names := [...]string{
        "Accountant",
        "Analysist",
        "Manager",
        "Manufacturing",
        "Programmer",
        "Inventory",
        "Sales",
        "SoftwareEnginner",
    }
    if j < 0 || int(j) >= len(names) {
        return "Unknown"
    }
    return names[j]
}