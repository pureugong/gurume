You have to use Custom Plugins section to manage dictionaries 

- Create a zip file with the following directory structure:
```sh
tree
.
|__ dictionaries
    |__ userdict_ko.txt

zip gurume-dictionaries -r dictionaries
```
- Login to elastic cloud and go to Custom Plugins section
- Click on Add Plugin
- Fill in the relevant details and for the section Plugin Type select A bundle containing a dictionary or script
- Click on Create Plugin

- Go back to the Custom plugins page and click on the new plugin you just added.
- Scroll to the bottom and upload the zip file created in first step.

Now you have to update your cluster so that its available to all the nodes. 
To do this follow the steps below:

- Click on Deployment
- Select your cluster/deployment form the page
- On the menu in the left click on Edit.
- Scroll to the section Elasticsearch plugins and settings on the page. 
- Click on Manage plugins and settings.
- From the expanded list select your bundle (located under Custom Plugins section in the expanded list).
- On the bottom of the page click on Save Changes
- Wait for the update activity to complete. Once completed you can now use stopwords.txt as below:

```json
// exmapping
{
    //...
    "type": "nori_tokenizer",
    "decompound_mode": "mixed",
    "user_dictionary": "userdict_ko.txt"
    // ...
}
```
