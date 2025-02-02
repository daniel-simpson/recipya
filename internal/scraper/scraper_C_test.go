package scraper

import (
	"github.com/reaper47/recipya/internal/models"
	"testing"
)

func TestScraper_C(t *testing.T) {
	testcases := []testcase{
		{
			name: "cafedelites.com",
			in:   "https://cafedelites.com/butter-chicken",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner"},
				CookTime:      "PT30M",
				Cuisine:       models.Cuisine{Value: "Indian"},
				DatePublished: "2019-01-21T19:09:20+00:00",
				Description: models.Description{
					Value: "Butter Chicken is creamy and easy to make right at home in one pan with simple ingredients! Full of incredible flavours, it rivals any Indian restaurant! Aromatic golden chicken pieces in an incredible creamy curry sauce, this Butter Chicken recipe is one of the best you will try!",
				},
				Keywords: models.Keywords{Values: "butter chicken"},
				Image: models.Image{
					Value: "https://cafedelites.com/wp-content/uploads/2019/01/Butter-Chicken-IMAGE-64.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"28 oz (800g) boneless and skinless chicken thighs or breasts (cut into bite-sized pieces)",
						"1/2 cup plain yogurt",
						"1 1/2 tablespoons minced garlic",
						"1 tablespoon minced ginger ((or finely grated))",
						"2 teaspoons garam masala",
						"1 teaspoon turmeric",
						"1 teaspoon ground cumin",
						"1 teaspoon red chili powder",
						"1 teaspoon of salt",
						"2 tablespoons olive oil",
						"2 tablespoons ghee ((or 1 tbs butter + 1 tbs oil))",
						"1 large onion, (sliced or chopped)",
						"1 1/2 tablespoons garlic, (minced)",
						"1 tablespoon ginger, (minced or finely grated)",
						"1 1/2 teaspoons ground cumin",
						"1 1/2 teaspoons garam masala",
						"1 teaspoon ground coriander",
						"14 oz (400 g) crushed tomatoes",
						"1 teaspoon red chili powder ((adjust to your taste preference))",
						"1 1/4 teaspoons salt ((or to taste))",
						"1 cup of heavy or thickened cream ((or evaporated milk to save calories))",
						"1 tablespoon sugar",
						"1/2 teaspoon kasoori methi ((or dried fenugreek leaves))",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a bowl, combine chicken with all of the ingredients for the chicken marinade; let marinate for 30 minutes to an hour (or overnight if time allows).",
						"Heat oil in a large skillet or pot over medium-high heat. When sizzling, add chicken pieces in batches of two or three, making sure not to crowd the pan. Fry until browned for only 3 minutes on each side. Set aside and keep warm. (You will finish cooking the chicken in the sauce.)",
						"Heat butter or ghee in the same pan. Fry the onions until they start to sweat (about 6 minutes) while scraping up any browned bits stuck on the bottom of the pan.",
						"Add garlic and ginger and sauté for 1 minute until fragrant, then add ground coriander, cumin and garam masala. Let cook for about 20 seconds until fragrant, while stirring occasionally.",
						"Add crushed tomatoes, chili powder and salt. Let simmer for about 10-15 minutes, stirring occasionally until sauce thickens and becomes a deep brown red colour.",
						"Remove from heat, scoop mixture into a blender and blend until smooth. You may need to add a couple tablespoons of water to help it blend (up to 1/4 cup). Work in batches depending on the size of your blender.",
						"Pour the puréed sauce back into the pan. Stir the cream, sugar and crushed kasoori methi (or fenugreek leaves) through the sauce. Add the chicken with juices back into the pan and cook for an additional 8-10 minutes until chicken is cooked through and the sauce is thick and bubbling.",
						"Garnish with chopped cilantro and serve with fresh, hot garlic butter rice and fresh homemade Naan bread!",
					},
				},
				Name: "Butter Chicken",
				NutritionSchema: models.NutritionSchema{
					Calories:      "580 kcal",
					Carbohydrates: "17 g",
					Cholesterol:   "250 mg",
					Fat:           "41 g",
					Fiber:         "3 g",
					Protein:       "36 g",
					SaturatedFat:  "19 g",
					Servings:      "1",
					Sodium:        "1601 mg",
					Sugar:         "8 g",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 5},
				URL:      "https://cafedelites.com/butter-chicken"},
		},
		{
			name: "castironketo.com",
			in:   "https://www.castironketo.net/blog/balsamic-mushrooms-with-herbed-veggie-mash/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				CookTime:      "PT30M",
				Cuisine:       models.Cuisine{Value: "Italian"},
				DatePublished: "2022-03-06T03:40:00+00:00",
				Description: models.Description{
					Value: "This easy low-carb dinner is perfect for plant-based eaters or anyone looking to add more veggies " +
						"to their diet!",
				},
				Image: models.Image{
					Value: "https://www.castironketo.net/wp-content/uploads/2022/03/" +
						"Balsamic-Mushrooms-with-Herbed-Veggie-Mash.jpg-p1ft8631n71uvu1giuf7n1i1b1hil-1-scaled.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"10 ounces cremini mushrooms",
						"2 tablespoons olive oil",
						"3 tablespoon water",
						"3 tablespoons balsamic vinegar",
						"¼ teaspoon sea salt",
						"¼ teaspoon freshly cracked black pepper",
						"3 tablespoons unsalted dairy-free butter (or regular if not dairy-free)",
						"2 cloves garlic (minced)",
						"2 tablespoons minced fresh herbs (we used rosemary, oregano, and thyme)",
						"4 cups cauliflower florets",
						"⅔ cup unsweetened plain almond milk (or heavy cream if not dairy-free)",
						"1 ½ cups chopped kale",
						"Sea salt and freshly cracked pepper (to taste)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"First, make the veggie mash by heating the butter in a large 10.5” skillet over medium-high heat. " +
							"Once hot add in the garlic and fresh herbs cook for 30 seconds until fragrant. Add in the " +
							"cauliflower and heavy cream. Cover and simmer for 15 minutes until the florets are soft. Transfer " +
							"the mixture to a food processor and blend until smooth. If needed, add an additional tablespoon or " +
							"two of heavy cream to reach your desired consistency. Season with salt and pepper to taste.",
						"To the same skillet over medium heat add the kale with 2 tablespoons of water. Cover and cook 3-4 minutes " +
							"until wilted. Stir the kale into the mashed cauliflower and divide between two bowls.",
						"In the same skillet over medium-high heat, heat the olive oil. Once hot add in the garlic and mushrooms. " +
							"Add 1 tablespoon water to the skillet, cook for 5-7 minutes until the mushrooms are soft. Add " +
							"in the balsamic vinegar, salt, and black pepper. Cook another 1-2 minutes until the vinegar has " +
							"reduced and is thick. Top the bowls with the mushrooms and serve.",
					},
				},
				Name: "Balsamic Mushrooms with Herbed Veggie Mash",
				NutritionSchema: models.NutritionSchema{
					Calories:       "201 kcal",
					Carbohydrates:  "12 g",
					Cholesterol:    "",
					Fat:            "16 g",
					Fiber:          "3 g",
					Protein:        "4 g",
					SaturatedFat:   "3 g",
					Servings:       "1",
					Sodium:         "310 mg",
					Sugar:          "5 g",
					TransFat:       "2 g",
					UnsaturatedFat: "13 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.castironketo.net/blog/balsamic-mushrooms-with-herbed-veggie-mash/"},
		},
		{
			name: "cdkitchen.com",
			in:   "https://www.cdkitchen.com/recipes/recs/285/MerleHaggardsRainbowStew65112.shtml",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "Merle Haggard's Rainbow Stew",
				Description: models.Description{
					Value: `This colorful stew named for American country singer, Merle Haggard's song "Rainbow Stew", is ` +
						"loaded with sausage, chicken, beans, and fresh vegetables.",
				},
				Yield:    models.Yield{Value: 6},
				CookTime: "PT1H20M",
				Category: models.Category{Value: "stews"},
				Ingredients: models.Ingredients{
					Values: []string{
						"5 tablespoons canola oil, divided",
						"1 pound kielbasa, chorizo or andouille sausage, cut into 1/2-inch cubes",
						"1 pound boneless skinless chicken breasts, cut into 1-inch cubes",
						"3 cups chicken broth",
						"3 tablespoons all-purpose flour",
						"1/2 cup chopped red bell peppers",
						"1/2 cup chopped yellow bell peppers",
						"1/2 cup chopped green bell peppers",
						"1/2 cup chopped purple onion",
						"1 cup peeled and diced carrots",
						"1/2 cup chopped celery",
						"2 cloves garlic, minced",
						"1 cup peeled and cubed jicama",
						"2 tablespoons chopped parsley or cilantro",
						"1 can (16 ounce size) dark red kidney beans, rinsed and drained",
						"1 bay leaf, crumbled",
						"1 teaspoon summer savory, crumbled",
						"5 teaspoons cayenne pepper",
						"salt and pepper, to taste",
						"1/2 cup chopped green onions",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Heat 2 tablespoons of the oil in a Dutch oven over medium heat. Add the sausage and cook, stirring " +
							"frequently, until browned. Remove with a slotted spoon and set aside.",
						"Add the chicken to the pan and cook, stirring frequently, until browned. Remove the chicken with a " +
							"slotted spoon and add to the sausage. Drain the oil from the pan and return the chicken and " +
							"sausage to the pan.",
						"Add the broth and bring it to a simmer. Let cook until the chicken is cooked through.",
						"In a large skillet over medium heat, combine the flour and remaining oil and cook, stirring constantly, " +
							"until smooth. Stir in the bell peppers, onion, carrots, celery, garlic, jicama, and parsley. " +
							"Cook, stirring frequently, for 10 minutes.",
						"Transfer the vegetable mixture to the Dutch oven. Add the kidney beans, bay leaf, savory, and cayenne. " +
							"Bring to a boil then reduce the heat to a simmer. Let cook, uncovered, for 45 minutes, stirring frequently.",
						"Season to taste with salt and pepper. Add the green onions and mix well. Serve over rice.",
					},
				},
				NutritionSchema: models.NutritionSchema{
					Calories:      "565 calories",
					Carbohydrates: "28 grams carbohydrates",
					Fat:           "34 grams fat",
					Protein:       "36 grams protein",
				},
				URL: "https://www.cdkitchen.com/recipes/recs/285/MerleHaggardsRainbowStew65112.shtml",
			},
		},
		{
			name: "chefkoch.de",
			in:   "https://www.chefkoch.de/rezepte/1064631211795001/Knusprige-Ofenkartoffeln.html",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Raffiniert & preiswert"},
				CookTime:      "P0DT0H40M",
				DatePublished: "2008-05-26",
				Description: models.Description{
					Value: "Knusprige Ofenkartoffeln. Über 231 Bewertungen und für raffiniert befunden. Mit ► Portionsrechner ► Kochbuch ► Video-Tipps! Jetzt entdecken und ausprobieren!",
				},
				Keywords: models.Keywords{
					Values: "Backen,Vegetarisch,Saucen,Dips,Beilage,raffiniert oder preiswert,einfach,Kartoffel,Snack",
				},
				Image: models.Image{
					Value: "https://img.chefkoch-cdn.de/rezepte/1064631211795001/bilder/1329056/crop-960x540/" +
						"knusprige-ofenkartoffeln.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"10 m.-große Kartoffeln, festkochende",
						"3 EL Olivenöl",
						"1 TL Meersalz , grobes, bei Bedarf auch mehr",
						"1 TL Paprikapulver, edelsüßes",
						"1 TL Currypulver",
						"etwas Chilipulver",
						"1 EL Rosmarin , getrocknet (frische Nadeln schmecken natürlich intensiver)",
						"etwas Pfeffer , aus der Mühle",
						"1 Becher saure Sahne",
						"1 Zitrone(n) , der Saft davon",
						"etwas Salz und Pfeffer",
						"1 EL Kräuter nach Wahl, italienische",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Die geschälten, geviertelten Kartoffeln gut mit Küchenkrepp abtrocknen. Die Kartoffeln in eine " +
							"große Schale geben, die Gewürze dazugeben und alles gut mit den Händen durchmengen. Nun das Öl " +
							"dazugeben und nochmals durchmengen. Die Kartoffeln auf ein Blech geben. Ich fette das Blech nicht " +
							"noch zusätzlich ein, da das Öl an den Kartoffeln bereits genügt.",
						"Im vorgeheizten Backofen bei 200 °C Ober-/Unterhitze auf mittlerer Schiene ca. 30 - 40 Min. (je " +
							"nach Größe der Kartoffelspalten) backen.",
						"Alle 10 Min. müssen die Spalten gewendet werden, damit sie von allen Seiten schön kross werden. " +
							"Wenn die Spalten anfangs etwas am Blech festbacken - nicht schlimm, einfach mit einem Pfannenwender lösen.",
						"Nun aus der sauren Sahne, dem Zitronensaft, den Gewürzen und den Kräutern einen Dip anrühren. Der " +
							"kann bei Bedarf natürlich auch mit etwas Knoblauch verfeinert werden.",
						"Dazu gibt es bei uns Fisch oder Hähnchenschenkel.",
					},
				},
				Name: "Knusprige Ofenkartoffeln",
				NutritionSchema: models.NutritionSchema{
					Calories:      "389 kcal",
					Carbohydrates: "53,93g",
					Fat:           "14,10g",
					Protein:       "8,52g",
					Servings:      "1",
				},
				PrepTime: "P0DT0H20M",
				Yield:    models.Yield{Value: 3},
				URL:      "https://www.chefkoch.de/rezepte/1064631211795001/Knusprige-Ofenkartoffeln.html",
			},
		},
		{
			name: "chefnini.com",
			in:   "https://www.chefnini.com/ramen-vegetarien/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Pâtes"},
				DateModified:  "2023-09-14T10:18:40+02:00",
				DatePublished: "2023-09-14T10:18:35+02:00",
				Description: models.Description{
					Value: "Depuis le temps que je voulais vous partager ma recette de ramen que je fais depuis 1 an maintenant. Au fur et à mesure, j’ai amélioré la recette pour la faire à notre goût. Ce qu’on aime le plus dans cette recette, ce sont les œufs mollets marinés. 🤤",
				},
				Keywords: models.Keywords{Values: "carotte,champignon,courgette,oeuf"},
				Image: models.Image{
					Value: "https://static.chefnini.com/wp-content/uploads/2023/09/ramen-vegetarien7.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 oeufs", "3 cuillères à soupe de sauce soja (salée ou shoyu)",
						"1/2 cuillère à soupe de vinaigre de riz", "2 cuillères à café de sucre",
						"140 g de nouilles ramen", "1 belle courgette", "2-3 carottes",
						"150 g de champignons de Paris", "1 oignon", "1 L d’eau",
						"1 cube de bouillon de légumes",
						"4 cuillères à soupe de sauce soja (salée ou shoyu)",
						"huile de sésame (facultatif)", "Ciboule ou ciboulette",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"1- 6h à l’avance (+ c’est mariné, meilleur c’est !) : faites cuire les œufs 6 minutes à l’eau bouillante pour les avoir mollet.",
						"2- Ecalez-les.",
						"3- Faites chauffer dans une petite casserole la sauce soja, le vinaigre de riz et le sucre.",
						"4- Versez dans un sachet congélation et ajoutez les œufs.",
						"5- Placez au frais pour 6 h en les retournant de temps en temps.",
						"1- Portez à ébullition l’eau avec la sauce soja et le cube de bouillon de légumes.",
						"2- Pendant ce temps, épluchez et émincez l’oignon.",
						"3- Faites-le cuire 10 minutes dans un peu de matière grasse (idéalement de l’huile de sésame).",
						"4- Lorsqu’il est cuit et un peu doré, ajoutez-le au bouillon. Poursuivez la cuisson à frémissement.",
						"5- Pendant ce temps, épluchez et émincez les champignons. Faites-les cuire dans la même poêle où ont cuits les oignons.",
						"6- Epluchez et coupez en julienne la courgette et les carottes.",
						"7- Faites cuire les ramens dans le bouillon selon le temps indiqué sur le paquet.",
						"8- Retirez les ramens et répartissez-les dans des assiettes creuses.",
						"9- A la place des ramens, faites cuire les légumes dans le bouillon. Laissez cuire 5 minutes. Les légumes doivent restés un peu croquants.",
						"10- Répartissez dans les ramens les légumes, les champignons, arrosez de bouillon.",
						"11- Ajoutez les œufs mollets coupés en deux.",
						"12- Si vous en avez, ajoutez de la ciboule ciselée ou de la ciboulette.",
					},
				},
				Name:  "Ramen végétarien",
				Yield: models.Yield{Value: 2},
				URL:   "https://www.chefnini.com/ramen-vegetarien/",
			},
		},
		{
			name: "chefsavvy.com",
			in:   "https://chefsavvy.com/crispy-baked-chicken-wings/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Appetizer"},
				CookTime:      "PT45M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2023-10-27T22:16:53+00:00",
				Description: models.Description{
					Value: "Whether drenched in smoky barbecue, zesty buffalo, or a sweet and sour glaze, these Crispy Baked Chicken Wings are succulent on the inside and crispy on the outside for the perfect wing experience! Don&#39;t be fooled, the oven can make these wings just as tasty as deep frying!",
				},
				Keywords: models.Keywords{Values: "Baked Chicken Wings"},
				Image:    models.Image{Value: "https://chefsavvy.com/wp-content/uploads/crispy-baked-chicken-wings.jpg"},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 1/2 lbs chicken wings", "1/2 tbsp baking powder", "1/2 tsp salt",
						"1/4 tsp pepper", "1/2 tsp garlic powder", "1/2 tsp onion powder",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 425 degrees. Line a baking sheet with a slipmat or aluminum foil. Top with a wire rack, set aside.",
						"Dab the chicken wings off with a paper towel to ensure they are completely dry.",
						"Add the chicken wings to a medium bowl. Sprinkle with baking powder, salt, pepper, garlic powder and onion powder. Toss to coat the chicken wings in the seasonings.Place the chicken wings onto the prepared baking sheet and bake for 45 minutes (depending on the size of the chicken wings)",
						"Turn the chicken wings every 20 minutes to ensure they get crispy on both sides.",
						"Serve the chicken wings as is with ranch dressing or blue cheese for dipping. You could also toss these with your favorite wings sauce or my Honey Chipotle Sauce!",
					},
				},
				Name: "Crispy Baked Chicken Wings",
				NutritionSchema: models.NutritionSchema{
					Calories:       "204 kcal",
					Cholesterol:    "71 mg",
					Fat:            "15 g",
					Protein:        "17 g",
					SaturatedFat:   "4 g",
					Servings:       "1",
					Sodium:         "67 mg",
					TransFat:       "0.2 g",
					UnsaturatedFat: "9 g",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://chefsavvy.com/crispy-baked-chicken-wings/",
			},
		},
		{
			name: "closetcooking.com",
			in:   "https://www.closetcooking.com/chipotle-roast-sweet-potatoes/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				CookTime:      "PT20M",
				DateModified:  "2023-11-12T11:20:22+00:00",
				DatePublished: "2023-11-16T06:00:54+00:00",
				Description:   models.Description{Value: "Tender roast sweet potatoes with smoky chipotle chili warmth!"},
				Keywords:      models.Keywords{Values: "Food,Gluten-free,Recipe,Side Dish,Vegetable,Vegetarian"},
				Image: models.Image{
					Value: "https://www.closetcooking.com/wp-content/uploads/2023/11/Chipotle-Roast-Sweet-Potatoes-1200-0699.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 pounds sweet potatoes, washed, optionally peeled, and cut into 1 inch cubes",
						"2 tablespoons oil", "1 teaspoon chipotle chili powder",
						"1 teaspoon paprika (smoked)", "1 teaspoon ground cumin",
						"1/2 teaspoon garlic powder", "1/2 teaspoon salt", "1/2 teaspoon pepper",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Toss the sweet potato in the oil, chili powder, paprika, cumin, garlic, salt and pepper, place on a baking sheet in a single layer and roast in a preheated 425F/220C oven until tender, about 20 minutes, flipping halfway through.",
					},
				},
				Name: "Chipotle Roast Sweet Potatoes",
				NutritionSchema: models.NutritionSchema{
					Calories:      "Calories 214",
					Carbohydrates: "Carbs 46g",
					Cholesterol:   "Cholesterol 0",
					Fat:           "Fat 7g",
					Fiber:         "Fiber 7g",
					Protein:       "Protein 3g",
					SaturatedFat:  "Saturated 0.5g",
					Sodium:        "Sodium 436mg",
					Sugar:         "Sugars 9g",
					TransFat:      "Trans 0",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.closetcooking.com/chipotle-roast-sweet-potatoes/",
			},
		},
		{
			name: "comidinhasdochef.com",
			in:   "https://comidinhasdochef.com/pudim-no-copinho-para-festa/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Sobremesa"},
				CookTime:      "PT5M",
				Cuisine:       models.Cuisine{Value: "Brasileira"},
				DatePublished: "2022-03-06T20:46:27-03:00",
				Description: models.Description{
					Value: "Você sabia que pudim pode ser um excelente doce para servir em festas? Pois, é! Praticamente todo mundo " +
						"ama um pudim, né? E,  por isso, muitas pessoas têm aprendido como fazer pudim no copinho para festa " +
						"que é um sucesso! E, se você trabalha nesta área ou adora preparar as festas da sua família, com " +
						"certeza você vai amar esta receita. O pudim fica delicioso, bem levinho, daquele jeito que todo mundo " +
						"ama e no ponto certo com esta receita, além disso, com o passo a passo que preparamos, você vai " +
						"conseguir montar todos os copinhos para servir nas festas e comemorações. Com certeza, todo mundo " +
						"vai amar e ainda querer repetir várias e várias vezes! Portanto, fique atento nesta receita de pudim " +
						"no copinho para festa, que pode ser uma solução para as comemorações e festas infantis, assim como " +
						"uma fonte de renda extra para você oferecer em seus serviços!",
				},
				Keywords: models.Keywords{Values: "Pudim no Copinho para Festa"},
				Image: models.Image{
					Value: "https://comidinhasdochef.com/wp-content/uploads/2022/03/Pudim-no-Copinho-para-Festa00.jpg",
				},
				Ingredients: models.Ingredients{Values: []string{
					"1 xícara (chá) de Água",
					"2 xícaras (chá) de açúcar",
					"500 ml de leite",
					"0 e 1/2 colheres (sopa) de essência de baunilha",
					"2 caixas de creme de leite",
					"2 caixas de leite condensado",
					"10 colheres (sopa) de Água quente",
					"2 pacotinhos de gelatina incolor sem sabor",
				},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Em uma panela e no fogo médio vá adicionando o açúcar aos poucos, mexendo sem parar até derreter completamente;",
						"Adicione a água e continue mexendo até dissolver todas as pelotas formadas pelo açúcar;",
						"Quando estiver em consistência de calda bem liquida desligue o fogo e aguarde esfriar;",
						"Unte os copinhos (50 ml) com um fio de óleo e distribua toda a calda entre eles (cerca de 1 colher de sopa " +
							"em cada) e reserve.",
						"Em uma tigela dissolva toda a gelatina incolor na água quente e em seguida reserve;",
						"No liquidificador coloque o leite, a essência de baunilha, o creme de leite e o leite condensado;",
						"Bata por 2-3 minutos e acrescente a gelatina dissolvida;",
						"Bata novamente por mais 1 minuto ou até ficar bem homogêneo;",
						"Distribua a mistura entre os copinhos com a calda e para finalizar leve a geladeira por pelo menos 2 horas;",
						"Retire da geladeira e pronto, já pode servir.",
					},
				},
				Name: "Pudim no Copinho para Festa",
				NutritionSchema: models.NutritionSchema{
					Calories:       "215 kcal",
					Carbohydrates:  "36 g",
					Fat:            "6.5 g",
					Fiber:          "0 g",
					Protein:        "3.6 g",
					SaturatedFat:   "3.7 g",
					Servings:       "50",
					Sodium:         "57 mg",
					TransFat:       "0 g",
					UnsaturatedFat: "1.6 g",
				},
				PrepTime: "PT20M",
				Tools:    models.Tools{Values: []string(nil)},
				Yield:    models.Yield{Value: 50},
				URL:      "https://comidinhasdochef.com/pudim-no-copinho-para-festa/"},
		},
		{
			name: "cookeatshare.com",
			in:   "https://cookeatshare.com/recipes/balinese-bbq-pork-roast-babi-guling-81003",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Kid Friendly"},
				CookTime:      "PT120M",
				Cuisine:       models.Cuisine{Value: "Indonesian"},
				DatePublished: "2009-07-05",
				Description: models.Description{
					Value: "My cousin Brett recently went to Bali for his honeymoon, so I decided to try the most famous Balinese " +
						"dish for July 4th this summer.  This pork came out soft and tasty like the meat you eat at a Luau, but " +
						"with a delicious complexity of flavors from the chiles, garlic, ginger, lemongrass and turmeric.  " +
						"The process of slicing the roast in the middle and stuffing it with the marinade infused the entire " +
						"roast with flavor.  Our guests raved...and were still going back for more even after dessert!  " +
						"Make sure to serve the Balinese yellow rice.  A wonderful combination!",
				},
				Image: models.Image{
					Value: "https://s3.amazonaws.com/grecipes/public/pictures/recipes/40375/balinese_pork.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 boneless pork shoulder roast, about 3 pounds (or bone in roast, about 3.5 pounds)",
						"4 large shallots or 2 small onions",
						"4-8 Thai chiles or 2-4 jalapenos",
						"4 cloves garlic, peeled",
						"2 Tbsp chopped fresh ginger1 Tbsp chopped fresh turmeric or 1/2 tsp ground turmeric",
						"1 Tbsp chopped fresh galangal, or 1 Tbsp more ginger",
						"3 stalks fresh lemongrass, trimmed and finely chopped or 3 large strips lemon zest",
						"1.5 tsp ground coriander",
						"1 tsp fresh ground black pepper",
						"2 Tbsp fresh lime juice (about 1 large lime)",
						"1 Tbsp firmly packed light brown sugar",
						"2 tsp salt",
						"3 Tbsp vegetable oil",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Combine the shallots, chiles, garlic, ginger, turmeric, galangal, lemongrass, coriander, pepper, lime " +
							"juice, sugar, salt, and 1 Tbsp vegetable oil in a mortar and pound to a smooth paste with the " +
							"pestle, or puree in a blender or food processor to a smooth paste.",
						"Heat 2 Tbsp of oil in a wok or skillet over medium heat. Add the spice paste and saute until fragrant, " +
							"about 5 minutes. Stir frequently to avoid splattering, and run your fan on high. Remove from " +
							"pan and let cool. The paste can be made up to a day in advance and refrigerated in an airtight container.",
						"Trim the roast of excess fat, if any. Using a sharp knife, make a deep slice in the center of roast, " +
							"starting and ending 3/4 inch from the ends, and cutting almost though to the other side of the " +
							"roast. You should have a nice pocket about 6 inches long.",
						"Fill the pocket with spice paste and tie the roast back together with kitchen twine or pin it with " +
							"metal skewers.  Spread the remaining paste all over the outside of the roast. If any spice mixture " +
							"remains, set it aside to add during the grilling process.",
						"If you have a rotisserie, this is probably the best way to cook the roast. Preheat your grill to high " +
							"and cook for about 1.5 hours.",
						"I used the indirect cooking method. Move the charcoals to either side of the place where you plan to " +
							"cook the roast, and place a drip pan in the middle. If you have a gas grill where the coals are not " +
							"movable, either turn off the middle burner or just put a drip pan made of aluminum foil directly over " +
							"the center burner. To make a foil drip pan, just tear 3 pieces of foil about 16 inches long, overlap them " +
							"so they are the width you want, then fold the edges up about 2 inches to form a make-shift drip pan.",
						"After setting up the grill, preheat it on high. Turn the grill down and place the roast over the drip " +
							"pan. Adjust heat or coals so the internal temperature rests at about 350 F. This will ensure nice " +
							"browning without burning, and should result in a cook time of about 2 hours. Turn the meat occasionally " +
							"during cooking so all sides get equally browned, and rub on additional spice mixture if any.",
						"Transfer roast to cutting board or platter and let stand for 10 minutes before removing strings and " +
							"cutting into thin slices to serve.",
					},
				},
				Name: "Balinese BBQ Pork Roast - Babi Guling",
				NutritionSchema: models.NutritionSchema{
					Calories:      "",
					Carbohydrates: "4.05 g",
					Cholesterol:   "0 g",
					Fat:           "6.92 g",
					Fiber:         "0.4 g",
					Protein:       "0.26 g",
					SaturatedFat:  "0.52 g",
					Servings:      "20 g",
					Sodium:        "777 g",
					Sugar:         "2.36 g",
					TransFat:      "0.18 g",
				},
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 1},
				URL:      "https://cookeatshare.com/recipes/balinese-bbq-pork-roast-babi-guling-81003"},
		},
		{
			name: "cookieandkate.com",
			in:   "https://cookieandkate.com/honey-butter-cornbread-recipe/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Baked goods"},
				CookTime:      "PT35M",
				CookingMethod: models.CookingMethod{Value: "Baked"},
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2020-03-24",
				Description: models.Description{
					Value: "Make the best homemade cornbread with this easy recipe! It's fluffy on the inside, crisp around the " +
						"edges, and full of delicious honey-butter flavor. Recipe yields one large skillet of cornbread " +
						"(see recipe notes for alternate baking vessel options).",
				},
				Keywords: models.Keywords{Values: "skillet cornbread"},
				Image: models.Image{
					Value: "https://cookieandkate.com/images/2020/03/skillet-cornbread-recipe-2-2-225x225.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1/2 cup (1 stick) unsalted butter",
						"1 1/2 cups cornmeal, medium-grind or finer*",
						"1 1/2 cups white whole wheat flour, regular whole wheat flour or all-purpose flour",
						"1 1/2 teaspoons fine sea salt",
						"2 teaspoons baking powder",
						"1/2 teaspoon baking soda",
						"3 large eggs, at room temperature**",
						"2/3 cup honey or maple syrup",
						"1 1/2 cups milk of choice (regular cow&#8217;s milk, almond or oat milk, etc.), at room temperature",
						"Optional serving suggestions: additional butter, honey or jam",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat the oven to 375 degrees Fahrenheit. Place the butter in a large (12-inch) cast iron skillet and " +
							"place the skillet in the oven to melt the butter, about 5 to 13 minutes (keep an extra eye on " +
							"it as time goes on—we want it to get bubbly and lightly browned, but not burnt).",
						"Meanwhile, in a large bowl, combine the cornmeal, flour, salt, baking powder and baking soda. Stir to " +
							"combine, and set aside. In a medium bowl, whisk together the eggs and honey until fully blended. " +
							"Add the milk and whisk until evenly combined.",
						"Pour the liquid into the dry mixture, and stir just until moistened through (we&#8217;ll stir it more " +
							"soon). When the butter is melted and golden, use oven mitts (the skillet is crazy hot!) to remove " +
							"the skillet from the oven, and give it a gentle swirl to lightly coat about an inch up the sides.",
						"Pour the melted butter into the batter and stir just until incorporated. Pour the batter back into the " +
							"hot skillet, using a spatula to scrape all of the batter from the bowl into the skillet.",
						"Carefully return the skillet to the oven and bake until the bread is brown around the edge, springy to " +
							"the touch, and a toothpick inserted in the center comes out clean with just a few crumbs, 25 to 30 " +
							"minutes. Carefully (with oven mitts), place the skillet on a cooling rack. Let it cool for at " +
							"least 5 minutes before slicing and serving—perhaps with extra butter, honey or jam on the side.",
						"This cornbread will keep at room temperature in a sealed container for up to 3 days, or up to a week in " +
							"the refrigerator. You can also freeze it for up to 3 months. Gently reheat before serving.",
					},
				},
				Name: "Honey Butter Cornbread",
				NutritionSchema: models.NutritionSchema{
					Calories:      "200 calories",
					Carbohydrates: "30.3 g",
					Cholesterol:   "52 mg",
					Fat:           "7.8 g",
					Fiber:         "2.1 g",
					Protein:       "4.5 g",
					SaturatedFat:  "4.3 g",
					Servings:      "1",
					Sodium:        "268.9 mg",
					Sugar:         "12.9 g",
					TransFat:      "0 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 16},
				URL:      "https://cookieandkate.com/honey-butter-cornbread-recipe/"},
		},
		{
			name: "cookpad.com",
			in:   "https://cookpad.com/recipe/4610651",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "本格バターチキンカレー"},
				CookTime:      "PT40M",
				DateModified:  "2018-02-19",
				DatePublished: "2018-02-19",
				Description: models.Description{
					Value: "本格バターチキンカレーのレシピです。 おうちにある材料で作れるバターチキンカレーです(*´`*)♡とっても簡単♬すぐ出来るのでぜひお試しください♡ʾʾ",
				},
				Image: models.Image{
					Value: "https://img.cpcdn.com/recipes/4610651/640x640c/6de3ac788480ce2787e5e39714ef0856?u=6992401&p=1519025894",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"♥鶏モモ肉 500g前後", "♥玉ねぎ 2個",
						"♥にんにくチューブ 5cm",
						"♥生姜チューブ 5cm(なくても♡)",
						"♥カレー粉 大さじ1と1/2", "♥バター 大さじ2+大さじ3(60g)",
						"＊トマト缶 1缶", "＊コンソメ 小さじ1",
						"＊塩 小さじ(1〜)2弱", "＊砂糖 小さじ2", "＊水 100ml",
						"＊ケチャップ 大さじ1", "♥生クリーム 100ml",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"鶏モモ肉 は一口大に、 玉ねぎ は薄切り(orみじん切り)にします♪",
						"フライパンに バター(大さじ2) を熱し、鶏肉 に 塩胡椒 をふり表面をこんがり焼きます♪",
						"お鍋に バター(大さじ3) にんにくチューブ 生姜チューブ 玉ねぎ を入れてあめ色になるまでじっくり炒めます♪",
						"カレー粉 を加えて弱火で3分くらい炒めます♪",
						"＊ と 鶏肉(油分も) を加えて沸騰したら火が通るまで(10分程)煮ます♪",
						"仕上げに 生クリーム を加えて混ぜ、温まったらすぐ火を止めます♪\n完成♡♡\n更に仕上げに生\xe3\x82\xafリームをトッピングしました♡",
						"子供ごはんはこんな感じの盛り付けに♡♥",
					},
				},
				Name:  "30分で簡単♡本格バターチキンカレー♡",
				URL:   "https://cookpad.com/recipe/4610651",
				Yield: models.Yield{Value: 1},
			},
		},
		{
			name: "cook-talk.com",
			in:   "https://cook-talk.com/?p=5476",
			want: models.RecipeSchema{
				AtContext: "https://schema.org",
				AtType:    models.SchemaType{Value: "Recipe"},
				Category: models.Category{
					Value: "Птица",
				},
				DatePublished: "2014-03-20T15:23:00+00:00",
				Description: models.Description{
					Value: "Очень, очень вкусный салат. Там не только индейка, там еще сливочный сыр и орехи, полноценный ужин практически.",
				},
				Image: models.Image{Value: "https://cook-talk.com/archive/wp-content/uploads/2016/11/1-55.jpg"},
				Ingredients: models.Ingredients{
					Values: []string{
						"400 г фельд-салата (он же маш, он же рапунцель, он же полевой салат)",
						"300 г запеченного мяса индейки",
						"200 г сливочного сыра (Филадельфия, Almetto, Лучина)",
						"250 г помидоров черри",
						"1 ст. л. оливкового масла",
						"2-3 ст. л. сока грейпфрута",
						"50 г грецких орехов",
						"мелко нарезанный шнитт-лук или черемша",
						"пара долек грейпфрута для сервировки",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Орехи поставить в духовку на 220 градусов на 8 минут, готовые орехи остудить.",
						"Листочки фельд-салата разложить на тарелке. Выложить сверху мясо индейки, нарезанное на кубики. Сливочный сыр нарезать на такого же размера кубики и обвалять в мелко нарезанном шнитт-луке или черемше, тоже выложить на салат. (У меня не было ни того, ни другого, поэтому я обваливала в зеленом луке).",
						"Выложить на салат половинки помидоров. Сбрызнуть салат грейпфрутовым соком, а потом оливковым маслом. Солить не нужно.",
						"Сверху посыпать поджаренными и крупно порезанными орехами и положить дольки грейпфрута. Подавать со свежим хлебом.",
						"Приятного аппетита.",
					},
				},
				Name: "Салат из индейки",
				URL:  "https://cook-talk.com/?p=5476",
			},
		},
		{
			name: "coop.se",
			in:   "https://www.coop.se/recept/appelkaka-med-havregryn/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				DateCreated:   "2019-03-21 10:14:50",
				DatePublished: "2019-03-21 10:14:50",
				Description: models.Description{
					Value: "Här hittar du ett recept på äppelkaka med havrygryn som är enkel att göra och så god. Mycket crunch och saftig!",
				},
				Ingredients: models.Ingredients{Values: []string{}},
				Name:        "Äppelkaka med havregryn",
				URL:         "https://www.coop.se/recept/appelkaka-med-havregryn/",
			},
		},
		{
			name: "copykat.com",
			in:   "https://copykat.com/mcdonalds-egg-mcmuffin",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Breakfast"},
				CookTime:      "PT5M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2023-02-16T05:11:12+00:00",
				Description: models.Description{
					Value: "Learn how to make a McDonalds Egg McMuffin at home with this easy copycat recipe. Find out the secret " +
						"to making perfect egg rings for a breakfast sandwich.",
				},
				Keywords: models.Keywords{Values: "McDonald's Egg McMuffin"},
				Image: models.Image{
					Value: "https://copykat.com/wp-content/uploads/2020/04/McDonalds-Egg-McMuffin-Pin3.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 tablespoons softened butter (butter has divided uses)",
						"4 English Muffins",
						"4 slices Canadian Bacon",
						"4 eggs",
						"1/2 cup water",
						"4 slices American cheese",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Split open English Muffins and place them into a toaster, toast the English Muffins.",
						"In a non-stick skillet over medium heat, cook Candian bacon on both sides for about 1 to 2 minutes in " +
							"two teaspoons of butter. The bacon should begin to just brown.",
						"While the Canadian bacon is cooking, remove the English muffins from the toaster and spread softened " +
							"butter on both halves.",
						"Place the 1 slice of Canadian bacon on each English Muffin bottom.",
						"Add about 1 tablespoon of butter to the same skillet where you cooked the bacon.",
						"Place the quart-sized canning lids screw size up (or you can use an egg ring) into the skillet.",
						"Spray the canning lid with non-stick spray. Crack an egg into each of the rings.",
						"Break the yolk with a fork. Pour about 1/2 cup of water into the skillet, and place a lid on top. Cook " +
							"until the eggs are set, it should take about two minutes.",
						"Gently remove the eggs from the rings, and place one egg on each piece of Canadian bacon.",
						"Top each egg with one slice of American cheese, top cheese with the top of the English muffin.",
						"Wrap each egg McMuffin with foil or parchment paper. Wait about 30 seconds before serving.",
					},
				},
				Name: "McDonald's Egg McMuffin",
				NutritionSchema: models.NutritionSchema{
					Calories:      "420 kcal",
					Carbohydrates: "28 g",
					Cholesterol:   "229 mg",
					Fat:           "25 g",
					Fiber:         "2 g",
					Protein:       "20 g",
					SaturatedFat:  "13 g",
					Servings:      "1",
					Sodium:        "1037 mg",
					Sugar:         "1 g",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://copykat.com/mcdonalds-egg-mcmuffin",
			},
		},
		{
			name: "costco.com",
			in:   "https://www.costco.com/connection-recipe-chicken-salad-grapes-walnuts-blue-cheese-march-2023.html",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{},
				CookTime:      "",
				CookingMethod: models.CookingMethod{},
				Cuisine:       models.Cuisine{},
				DateCreated:   "",
				DateModified:  "",
				DatePublished: "",
				Description:   models.Description{},
				Keywords:      models.Keywords{},
				Image:         models.Image{Value: "https://mobilecontent.costco.com/live/resource/img/static-us-connection-march-23/03_23_FTT_ChickenSaladRedGrapesWalnutsBlueCheese.jpg"},
				Ingredients: models.Ingredients{
					Values: []string{
						"6 cups rotisserie chicken, shredded",
						"3 cups red seedless grapes, halved lengthwise",
						"2 celery ribs (about 1 cup), thinly sliced",
						"1 cup walnuts, toasted and chopped",
						"½ cup blue cheese, crumbled",
						"2 cups mayonnaise",
						"2 Tbsp sherry vinegar",
						"1 Tbsp fresh thyme leaves",
						"2 tsp lemon zest",
						"½ tsp garlic powder",
						"½ tsp kosher salt, or to taste",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Combine chicken, grapes, celery, walnuts and blue cheese in a large mixing bowl.",
						"In a medium-size mixing bowl, blend mayonnaise, sherry vinegar, thyme, lemon zest, garlic powder and salt.",
						"Fold the dressing into the chicken-grape mixture and combine well. Adjust seasonings as desired. Serve in Bibb lettuce cups or as a sandwich filling. Makes 6 servings.",
					},
				},
				Name:     "Chicken Salad with Red Grapes, Walnuts and Blue Cheese",
				PrepTime: "",
				Yield:    models.Yield{},
				URL:      "https://www.costco.com/connection-recipe-chicken-salad-grapes-walnuts-blue-cheese-march-2023.html",
			},
		},
		{
			name: "countryliving.com",
			in:   "https://www.countryliving.com/food-drinks/a39298988/braised-turkey-wings-recipe/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Sunday lunch"},
				CookTime:      "PT0S",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-03-04T00:01:33.158861Z EST",
				Description: models.Description{
					Value: "This Southern dish is just bursting with flavor.",
				},
				Keywords: models.Keywords{
					Values: "American, Southern, Sunday lunch, comfort food, dinner",
				},
				Image: models.Image{
					Value: "https://hips.hearstapps.com/hmg-prod/images/braised-turkey-wings-clx040122-1646247632.jpg?crop=0.878xw:0.585xh;0,0.223xh&resize=1200:*",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"3 whole turkey wings (about 3 pounds total)",
						"Kosher salt and freshly ground black pepper",
						"2 tbsp. olive oil",
						"1 c. chopped yellow onion",
						"1/2 c. chopped carrots",
						"1/2 c. chopped celery",
						"2 cloves garlic, chopped",
						"2 tsp. chopped fresh rosemary",
						"2 tsp. chopped fresh sage",
						"1 tsp. chopped fresh thyme",
						"2 bay leaves",
						"2 tbsp. all-purpose flour",
						"4 c. turkey or chicken stock",
						"Cooked white rice, for serving",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 350°F. Season wings with salt and pepper. Heat oil in a large Dutch oven over " +
							"medium-high heat. Add wings and cook, turning once, until golden brown, 4 to 5 minutes. " +
							"Transfer to a plate; reserve pot.",
						"Reduce heat to medium. Add onion, carrots, and celery to reserved pot. Cook, stirring occasionally, " +
							"until onion is translucent, 6 to 8 minutes. Add garlic, rosemary, sage, thyme, and bay leaves. " +
							"Cook, stirring, until garlic is fragrant, about 1 minute. Sprinkle in flour and cook, stirring, " +
							"until flour becomes a medium brown shade (like the color of caramel), 4 to 5 minutes. While " +
							"stirring, slowly pour in half of stock. Return wings to pot and pour in remaining stock until " +
							"wings are 2/3 covered by liquid. Cover and bake until wings are tender, 2 to 2 1/2 hours. Serve over rice.",
					},
				},
				Name: "Braised Turkey Wings",

				PrepTime: "PT40M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.countryliving.com/food-drinks/a39298988/braised-turkey-wings-recipe/",
			},
		},
		{
			name: "creativecanning.com",
			in:   "https://creativecanning.com/caramelized-onion-jam/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				CookTime:      "PT1H15M",
				DatePublished: "2022-09-05",
				Description: models.Description{
					Value: "Caramelized onion jam is a unique savory jam that&#x27;s delicious no matter where you use it.",
				},
				Image: models.Image{
					Value: "https://creativecanning.com/wp-content/uploads/2022/09/Caramelized-Onion-Jam-31-720x720.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 Tablespoon cooking oil, neutral, such as sunflower or grape seed",
						"3 pounds red onions, sliced",
						"1 cup maple sugar (note this is different from granulated sugar)",
						"2 Tablespoons bottled lemon juice", "1½ cups cider vinegar",
						"2 teaspoons kosher salt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"If canning, prepare a water bath canner, jars and lids before beginning. If not canning, no worries, just let it cool on the counter in jars before storing in the refrigerator.",
						"Heat oil in a heavy-bottomed pan or Dutch oven over medium-high heat, then add the sliced onions and stir to coat.",
						"Reduce the heat to medium and cook the onions for 45-50 minutes, stirring often.",
						"Prepare jars and a hot water bath.",
						"Add maple sugar to the cooked onions and stir to dissolve the sugar.",
						"Increase the heat to medium-high and cook the onions and sugar for 3-4 minutes, stirring constantly.",
						"Carefully stir in the vinegar and lemon juice, being cautious of splattering and vinegar vapor.",
						"Reduce the heat to medium and add salt to the mixture.",
						"Cook the onion mixture until the vinegar has reduced to a sticky syrup, about 10-12 minutes.",
						`Pour or ladle the hot jam into prepared hot jars, leaving ½ " headspace at the top of the jars.`,
						"Wipe the rims of the jars with a clean cloth, and apply the 2-part canning lids.",
						"Preserve the homemade onion jam using a hot water bath canning method for 10 minutes, adjusting time if necessary for altitude.",
						"After processing, turn off the heat and leave jars in the canner, uncovered, for 5 minutes. Remove the jars from the canner and let them cool for 12-24 hours, then check that the jars have sealed. Refrigerate any unsealed jars.",
						"Sealed red onion jam will retain peak quality for up to 18 months stored in a cool, dry place or pantry. Refrigerate the jam once opened.",
					},
				},
				Name:     "Caramelized Onion Jam",
				PrepTime: "PT20M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://creativecanning.com/caramelized-onion-jam/",
			},
		},
		{
			name: "cucchiaio.it",
			in:   "https://www.cucchiaio.it/ricetta/baccala-in-crosta-senza-glutine/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Secondi"},
				DateModified:  "2023-11-17",
				DatePublished: "2023-11-17",
				Description: models.Description{
					Value: "Il baccalà in crosta&nbsp;senza glutine è un raffinato secondo di pesce ideale per le cene dei giorni di festa. Sotto uno strato di fragrante pasta sfoglia si nasconde un tenero filetto di baccalà, avvolto in un battuto a base di pomodori, olive e capperi che rende il risultato finale saporito e goloso.",
				},
				Keywords: models.Keywords{Values: "Baccalà in crosta senza glutine,al forno"},
				Image: models.Image{
					Value: "https://www.cucchiaio.it/content/cucchiaio/it/ricette/2021/11/baccala-in-crosta-senza-glutine/jcr:content/header-par/image-single.img10.jpg/1671010226009.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 filetto intero di baccalà già ammollato (di 600-700 g circa)",
						"150 g di pomodori secchi sott’olio", "2 cucchiai di capperi sott’aceto",
						"2 cucchiai di olive taggiasche denocciolate",
						"1-2 rotoli di pasta sfoglia senza glutine", "1 tuorlo", "1 goccio di latte",
						"1 cucchiaino di semi di papavero", "1 ciuffo di prezzemolo", "sale", "pepe",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Per preparare il baccalà in crosta senza glutine iniziate a deliscare il filetto di baccalà con una pinzetta, poi con un coltello ben affilato rimuovete la pelle del pesce (potete anche lasciarla, se preferite). Scolate i pomodori, le olive e i capperi e tritateli al coltello cercando di ottenere un battuto non troppo fine. Riunite il battuto in una ciotola, aggiungete il prezzemolo tritato e mescolate.",
						"Sistemate la pasta sfoglia rettangolare in maniera orizzontale di fronte a voi su un foglio di carta forno. Spalmate soltanto la parte centrale con un po’ di battuto mediterraneo. Tamponate il baccalà con carta assorbente e sistematelo al centro della pasta sfoglia, esattamente sulla parte cosparsa di battuto. Pepate poi distribuite il resto del battuto sul baccalà.",
						"Coprite il baccalà con un secondo foglio di pasta sfoglia, poi con un coltello affilato ritagliate le abbondanze laterali. Usate i ritagli di pasta per decorare il guscio di sfoglia. Spennellate tutto con il tuorlo sbattuto con il latte, cospargete con i semi di papavero e cuocete in forno statico preriscaldato a 190° per circa 25 minuti, fino a doratura.",
						"Sfornate il baccalà in crosta senza glutine, lasciate intiepidire per almeno 15 minuti prima di servirlo tagliato a tranci.",
					},
				},
				Name:  "Baccalà in crosta senza glutine",
				URL:   "https://www.cucchiaio.it/ricetta/baccala-in-crosta-senza-glutine/",
				Yield: models.Yield{Value: 1},
			},
		},
		{
			name: "cuisineaz.com",
			in:   "https://www.cuisineaz.com/recettes/champignons-farcis-au-fromage-brie-87449.aspx",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Legumes Farcis"},
				CookTime:      "PT15M",
				Cuisine:       models.Cuisine{Value: "French"},
				DatePublished: "2016-06-06T14:39:27+02:00",
				Image: models.Image{
					Value: "https://img.cuisineaz.com/660x660/2016/06/06/i75661-champignons-farcis-au-fromage-brie.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"16 Champignon(s) de paris",
						"0.5 Brie",
						"1 Échalote(s)",
						"1 c. à soupe Crème fraîche",
						"1 Tranche(s) de jambon blanc",
						"2 c. à soupe Chapelure",
						"Sel poivre",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Etape 1",
						"Préchauffez le four à 180°C.",
						"Etape 2",
						"Enlevez le pied des champignons de Paris et nettoyez les chapeaux.",
						"Etape 3",
						"Epluchez l'échalote et coupez-la en quatre.",
						"Etape 4", "Coupez le fromage Brie en gros morceaux.",
						"Etape 5",
						"Dans le bol d'un robot mixeur, placez les morceaux d'échalote et de fromage Brie, la crème fraîche, la tranche de jambon blanc, la chapelure, du sel et du poivre. Mixez jusqu’à obtenir une crème bien lisse et homogène.",
						"Etape 6",
						"Répartissez la crème dans les champignons et disposez-les sur une plaque du four recouverte de papier sulfurisé.",
						"Etape 7",
						"Enfournez pendant 10 à 15 minutes.",
						"Etape 8",
						"Servez immédiatement accompagné de volaille.",
					},
				},
				Name:     "Champignons farcis au fromage Brie",
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.cuisineaz.com/recettes/champignons-farcis-au-fromage-brie-87449.aspx",
			},
		},
		{
			name: "cybercook.com.br",
			in: "https://cybercook.com.br/receitas/peixes-e-frutos-do-mar/receita-de-file-de-tilapia-com-batatas-82273?" +
				"receita-do-dia",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Category:  models.Category{Value: "Peixes e Frutos do Mar"},
				CookTime:  "PT1H",
				Description: models.Description{
					Value: "Já experimentou essa deliciosa receita de Filé de Tilápia com Batatas? No CyberCook você encontra " +
						"essa e outras receitas. Saiba mais!",
				},
				Image: models.Image{
					Value: "https://img.cybercook.com.br/receitas/273/file-de-tilapia-com-batatas.jpeg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"Tilápia 500 gramas",
						"Batata 3 unidades",
						"Requeijão 1 copo",
						"Sal a gosto",
						"Azeite 1 colher (sopa)",
						"Cebola 1 unidade",
						"Salsinha 1 colher (sopa)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Cozinhe as batatas al dente.",
						"Unte um refratário com azeite e requeijão, coloque as batatas e as cebolas, em cima coloque o peixe " +
							"temperado com sal e pimenta a gosto, e para finalizar acrescente o requeijão, leve ao " +
							"forno por Aproximadamente 40 minutos ou até dourar.",
						"Decore com salsinhas picadinhas",
						"Sirva a seguir.",
					},
				},
				Name: "Filé de Tilápia com Batatas",
				NutritionSchema: models.NutritionSchema{
					Calories: "274.20",
				},
				PrepTime: "PT1H",
				Tools:    models.Tools{Values: []string(nil)},
				Yield:    models.Yield{Value: 5},
				URL:      "https://cybercook.com.br/receitas/peixes-e-frutos-do-mar/receita-de-file-de-tilapia-com-batatas-82273",
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			test(t, tc)
		})
	}
}
