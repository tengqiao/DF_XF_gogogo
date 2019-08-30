import requests
import json

def getImagesNames(repo_ip, repo_port):
    url = "http://" + repo_ip + ":" + repo_port + "/v2/_catalog"
    req = requests.get(url).content.strip()
    req_dic = json.loads(req)
#get images    
    images_list = req_dic.get("repositories")
    for images_name in images_list:
        url_tags = "http://" + repo_ip + ":" + repo_port + "/v2/" +  images_name + "/tags/list"
        req_tags = requests.get(url_tags).content.strip()
        req_tags_dic = json.loads(req_tags)
        for tag in req_tags_dic.get("tags"):
            print repo_ip + ":" + repo_port + "/" + req_tags_dic.get("name") + ":" + tag

if __name__ == '__main__':
    getImagesNames("127.0.0.1", "5000")
