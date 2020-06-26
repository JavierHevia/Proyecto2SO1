from flask import Flask
from pymongo import MongoClient
import json

client = MongoClient("mongodb://admin:admin@cluster0-shard-00-00-k6sn1.mongodb.net:27017,cluster0-shard-00-$
#mycol = client["sopes"]
db = client.Base1.sopes

# the all-important app variable:
app = Flask(__name__)

@app.route("/")
def hello():
    return "Hello World "

@app.route("/datos/<json2>")
def article(json2):
   # company_type = {"name": "Javier", "municipio" : "SMP"}
   # result = db.insert_one(company_type)
   # result = db.insert_one(json2)
   xmess = json.loads(json2)
   print ("datos: ", xmess)
   # print ("Tamaño: ", len(xmess) )
   # tamanio = len(xmess)-1

   for x in xmess:
    #print("Name: ",x['name'])
    db.insert_one(x)
    print("Insertado..")

   # print ("=D ",xmess['name'])
   # print ("Resultado DB: ", result)
   respuesta = "Datos: " + json2
   return respuesta

if __name__ == "__main__":
        app.config['TEMPLATES_AUTO_RELOAD'] = True
        app.run(host='0.0.0.0', debug=True, port=8080)
