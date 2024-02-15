# pnc-tool (pefman no commentary tool)

## Description
the idea for this project was to make a tool to simplify the manipulation of tags for youtube videos.
I choose to make it open-source so that other can help add features and make the tool usefull.

## Roadmap

* export/import tags for specific video(s) 
* cron compability

### Installing
download the latest release and create config.json with following

``` 
{
    "APIKey": "xxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "ChannelId": "xxxxxxxxxxxxxxxxxxxxxxxxx"
} 
```

execute
```
./pnc-tool -gettags _xzfYzYHMn4
```
## License

This project is licensed under the Apache 2.0 License - see the LICENSE.md file for details