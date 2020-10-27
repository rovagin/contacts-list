FROM mongo:latest

COPY ./init.json /init.json
CMD mongoimport --collection=users_contacts --type=json --file=/init.json --jsonArray --uri=mongodb://admin:password@mongodb/users_contacts?authSource=admin