# Shopify-Challenge
Shopify Backend Developer Intern Challenge\
This is a program attempts to solve the problem outlined [here](https://docs.google.com/document/d/1PoxpoaJymXmFB3iCMhGL6js-ibht7GO_DkCF2elCySU/edit) for the fall 2022 internship position

#Brief introduction to this program
This program is written in Go, it runs as the backend system of an inventory that supports\
CRUD features. When delete an item, you can add comments (reason) to it, and also undelete the item

#How to run it!
Make sure you have Go 1.5+ installed, you can install it [here](https://go.dev/doc/install) \
To test the application, install postman, you can install it [here](https://www.postman.com/downloads/) \
Use command "git clone https://github.com/IsaacXie-star/Shopify-Challenge.git" to clone the code to your local computer\
To run the code, use command "go run main.go"\
The server will be listening at 127.0.0.1:8080

#Explanation of all the http interface provided by this server
1. /shop/inventory/item/get_item\
This is an HTTP GET method, this interface allows users to view a list of items\
variables should be added in the url for selecting purposes\
For example, /shop/inventory/item/get_item/?name=Apple&min_price=20
#
2. /shop/inventory/item/get_item_details
This is an HTTP GET method, this interface allows users to view the details of an item given a specific item id\
variables should be added in the url for selecting purposes\
For example, /shop/inventory/item/get_item_details/?id=3
#
3. /shop/inventory/item/edit_item
This is an HTTP POST method, this interface allows users to edit an item given a specific item id\
variables should be added in the Request body and in a post form, you can use Postman to test it\
This interface supports editing item's name, category, description, and price
For example, /shop/inventory/item/edit_item
post form: key:id value:1, key:price value:25.5
#
4. /shop/inventory/item/add_item
This is an HTTP POST method, this interface allows users to add an item\
variables should be added in the Request body and in a post form, you can use Postman to test it\
name, price, category are required, description is optional
For example, /shop/inventory/item/add_item
post form: key:price value:25.5, key:name value:apple, key:category value:fruit
#
5. /shop/inventory/item/delete_item
This is an HTTP POST method, this interface allows users to delete an item given a specific item id\
and also add a deletion comment. Variables should be added in the Request body and in a post form, \
you can use Postman to test it\
For example, /shop/inventory/item/delete_item
post form: key:id value 1, key:deletion_comment value:fot test
#
6. /shop/inventory/item/undelete_item
This is an HTTP POST method, this interface allows users to undelete a deleted item given a specific item id\
Variables should be added in the Request body and in a post form, you can use Postman to test it\
For example, /shop/inventory/item/undelete_item
post form: key:id value 1