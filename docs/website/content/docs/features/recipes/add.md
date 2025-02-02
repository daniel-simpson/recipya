---
title: Add
weight: 1
prev: /docs/features/recipes
---

You can add a recipe by clicking the blue **Add Recipe** in the middle of the application bar.

![](images/.webp)

## Adding Recipes

You will be presented with four different ways of adding a recipe to your collection.

- [Manual](#manual)
- [Scan](#scan)
- [Website](#website)
- [Import](#import)

![](images/add-recipe-options.webp)

### Manual

The simplest method of inputting a recipe involves completing a form that outlines your dish.
Mandatory fields are indicated with an asterisk (*).

![](images/add-recipe-manual.webp)

You might find these shortcuts useful if you are filling the form from your computer.

| Shortcut     | How to enable                     | Result              |
|--------------|-----------------------------------|---------------------|
| Enter        | Focus on an ingredient text input | Add ingredient row  |
| Ctrl + Enter | Focus on an instruction text area | Add instruction row |

You can also reorder the ingredients and the instructions by dragging the arrow vertically.

### Scan

You can upload an image or a scan of a handwritten or printed recipe to add it to your collection. 
This option is useful for digitizing your and your family's paper recipes.

To do so, click the *Upload* button and select an image in your computer.

![](images/add-recipe-scan.webp)

{{< callout type="warning" >}}
You must have an [Azure AI Vision](/guide/docs/installation/system-requirements#azure-ai-vision) account to use this feature.
{{< /callout >}}

### Website

You can import any recipe from the supported websites. To do so, click the **Fetch** button, 
paste the recipe's URL or URLs, each on a new line, and click *Submit*. A notification will 
appear when the operation finished. 

Its action button performs an action based on the outcome. When you imported one URL and the 
operation succeeds, the action will redirect you to the recipe. Otherwise, you will be redirected
to the latest report.

To view all supported websites, please click on the <ins>supported</ins> word. You can scrape a 
website not in the supported list, but recipe extraction may fail. If it does, you can request 
support for the website by clicking the button that appears.

![](images/add-recipe-website.gif)

### Import

You can import recipes in the following formats:
- `.json`: If they adhere to the [Recipe schema](https://schema.org/Recipe) standard
- `.mxp`: Exported recipes from [MasterCook](https://www.mastercook.com)  
- `.txt`

![](images/add-recipe-import.webp)

You can upload either a single file or a zip archive containing multiple recipes.
The recipes in a zip file may be organized by folder. Each folder may contain the `.json` recipe file and an image 
file. All other files in a folder will be ignored during processing. Here is an 
[example](https://sea.musicavis.ca/f/683b9b9a7cc84e1bac0c/?dl=1) of how such zip may look like.
