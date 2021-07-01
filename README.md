# What is Chasqi?

##### Is your Backend ready for production? 
##### <br/>Do you have a stable infrastructure? 
##### <br/>Are you about to launch your MVP?

If any of these qustions concern you then you might probably have a look at Chasqi.
Chasqi is a load-testing tool that helps you to immitate real world traffic. It supports configuration through YAML,
where you can specify the REST API endpoints of your backend that Chsqi will call with specified parameters or response values
returned from previously visited endpoints. Chasqi comes with pre-defined traffic load settings, which lets you see how your 
system behaves when it receives traffic if you were 
* Twitter
* A mid-sized startup (around 50 req / s)
* A startup that just launched and receives an unpredicted amount of user engagement (e.g. if your startup becomes successful over night)

Unless you haven't launched yet, it is recommended to not use Chasi on your production environment as it could in worst case
take down your system or blow up your AWS bill (if you're on AWS). Use it instead on your dev environment. 

## Navigation
You configure chasqi by setting up a yaml file. See demo.yaml as reference. Once you've done that you can call 
`chasqi -f my_yaml_file.yaml` and chasqi will start testing your infrastructure.

## Logging
There is no logging yet