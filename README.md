# jellyfin_exporter
This is a Prometheus (https://prometheus.io) metrics exporter for Jellyfin (https://jellyfin.org).

## Configuration
This exporter is configured via environment variables:
|      Name          |      Example                       |         Explanation                |               
| -------------------|------------------------------------|------------------------------------|                                      
| JELLYFIN_APIURL    | http://demo.jellyfin.org           |Base APIURL of the Jellyfin Instance|             
| JELLYFIN_APIKEY    | 9e49ae09128847ee667cfhj367811efv   |Authentication Token                |
| WEB_LISTEN_ADDRESS |            :9249                   |Port of jellyfin_exporter           |

Or from Command Line Flags:
 * `--jellyfin.apiurl=http://demo.jellyfin.org`
 * `--jellyfin.apikey=9e49ae09128847ee667cfhj367811efv`
 * `--web.listen-address=:9249`
## Exported Metrics
General metrics:
* jellyfin_active_users                  
* jellyfin_active_streams_count           
* jellyfin_active_streams_direct_play_count 
* jellyfin_active_streams_transcode_count  
* jellyfin_movie_count      
* jellyfin_series_count     
* jellyfin_episode_count    
* jellyfin_artist_count     
* jellyfin_program_count    
* jellyfin_trailer_count    
* jellyfin_song_count       
* jellyfin_album_count      
* jellyfin_music_video_count 
* jellyfin_box_set_count     
* jellyfin_book_count       
* jellyfin_item_count       