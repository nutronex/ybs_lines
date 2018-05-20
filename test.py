#!/usr/bin/python3
import os,sys,json

def main():

    try:
        f = open(os.getcwd()+os.sep+"data"+os.sep+"bus-stop-data-by-id.js")
        bustops = json.load(f)
        f.close()
        bustop_ids = [x for x in bustops]

        f2=open(os.getcwd()+os.sep+"data"+os.sep+"bus-stop-ids-of-lines.js")
        buslines=json.load(f2)
        busline_ids=[x for x in buslines]
        f2.close()
    except Exception :
        print("files not found :p")
        sys.exit(0)

#data={}
#dist_merge = lambda x,y,z : data.get(y,None) and data[y]["l"].append(z) or data.update({str(x):{"t":j[x]["township"],"l":[z]}})
#bustop_listxx =     { dist_merge(x,j[x]["township"],y)  for x in j for y in j2 if  x in j2[y]}

    bustop_listx={}
    for x in bustop_ids:
        data =  { str(x): {"t":bustops[x]["township"]}}
        foo = []
        for y in busline_ids:
            if x in buslines[y]:
                foo.append(y)
        data[x].update({"l":foo})
        bustop_listx.update(data)

    townships={}
    for i in bustop_listx:
        tship = bustop_listx[i]["t"]
        line = bustop_listx[i]["l"]
        if not tship in townships:
            townships.update({tship:set(line)})
        else:
            townships[tship].update(line)
    for i in townships:print(i+str(townships[i]));print("-------------")

if __name__=="__main__":
    main()
