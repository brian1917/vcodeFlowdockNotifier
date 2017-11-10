# Veracode Flowdock Notifier Utility

## Description
Utility designed to be run in a build process after a Veracode scan to notify Flowdock that the scan completed. Optionally, the notification can also include the compliance policy assigned to that app and whether or not it's passing. For builds that don't wait for the Veracode scan to complete, the utility can be set to run on a schedule to provide notifications.

## Executables
Executables for Windows, Mac, and Linux will be available in the releases section of the repository (https://github.com/brian1917/vcodeFlowdockNotifier/releases)

## Running the Utility
The utility takes one argument - the location of the JSON config file. Run the utility as a command line
action at the end of the build (after Veracode completed):
`vCodeFlowdockNotifer appABCconfig.json`

## Configuration File
A sample config file is shown below. The 5 parameters below are require to be present.
```
{
    "credsFile": "/Users/bpitta/.veracode/credentials",     // Location of Veracode credentials file
    "appID": "123456",                                      // App ID in Veracode being targeted
    "flowdockToken": "530b7d5e03f51d4835dcfa2a4f248096",    // API token from Flowdock (this example is random)
    "flowdockOrg": "organization",                          // Name of organization used in Flowfock endpoint   
    "flowdockFlow": "flow",                                 // Name of Flowdock flow in endpoint      
    "onlyNotifyOnNotPass" : false,                          // Will only notify if the scan violates policy
    "includePolicyStatus" : true                            // Will include policy name and compliance status in notification
}
```
In this example, the utility will provide a notifcation for App ID 123456 in the Veracode Platform for all builds (not just ones that don't pass), and will
include policy information in the notification.

## Veracode Credentials File
Must be structured like the following:
```
[DEFAULT]
veracode_api_key_id = ID HERE
veracode_api_key_secret = SECRET HERE
```

## Third-party Packages
github.com/brian1917/vcodeapi (https://godoc.org/github.com/brian1917/vcodeapi)