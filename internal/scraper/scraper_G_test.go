package scraper

import (
	"github.com/reaper47/recipya/internal/models"
	"testing"
)

func TestScraper_G(t *testing.T) {
	testcases := []testcase{
		{
			name: "gesund-aktiv.com",
			in:   "https://www.gesund-aktiv.com/rezepte/suppe/quitten-pastinaken-suppe",
			want: models.RecipeSchema{
				AtContext: "https://schema.org",
				AtType:    models.SchemaType{Value: "Recipe"},
				Description: models.Description{
					Value: "Die Antwort auf Herbstblues und Schmuddelwetter: Wärmender Genuss aus saisonalen Quitten, Pastinaken und milder Kokosmilch.",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 Stück Quitten", "2 Stück Pastinaken", "1 Stück Möhre",
						"2 Scheiben Ingwer", "1/2 Stück Zwiebel", "150 ml Kokosmilch",
						"200 ml Gemüsebrühe", "2 Esslöffel Sahne", "2 Esslöffel Rapsöl",
						"1 Prise Salz", "1 Prise Pfeffer",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Zu Beginn die Pastinaken, die Möhren und die halbe Zwiebel schälen und in kleine Würfel schneiden. Dann die Quitten mit einem Küchentuch abreiben, kurz abwaschen und ebenfalls schälen. Anschließend die Quitten vierteln, das Kerngehäuse großzügig entfernen und das Fruchtfleisch klein schneiden.",
						"Nun etwas Öl in eine erhitzte Pfanne geben und die Quittenwürfel einige Minuten leicht andünsten. Im Anschluss kommen die restlichen Gemüsewürfel und die Ingwer-Scheiben dazu. Alles zusammen wird acht bis zehn Minuten gedünstet und mit Salz und Pfeffer gewürzt.",
						"Das gedünstete Gemüse wird aus der Pfanne in einen höheren Topf umgefüllt. Jetzt kommen Gemüsebrühe und Kokosmilch hinzu und alles wird etwa 20 Minuten bei mittlerer Hitze leicht köchelnd ab und an umgerührt.",
						"Nachdem die Suppe gut durchgezogen ist, wird alles mit einem Stabmixer fein püriert. Zum Servieren dann noch mit etwas Sahne verfeinern.",
					},
				},
				Name: "Quitten-Pastinaken-Suppe",
				URL:  "https://www.gesund-aktiv.com/rezepte/suppe/quitten-pastinaken-suppe",
			},
		},
		{
			name: "giallozafferano.com",
			in:   "https://www.giallozafferano.com/recipes/Christmas-spice-cookies.html",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Sweets and desserts"},
				CookTime:      "PT15M",
				DateModified:  "2022-12-10 00:00:00",
				DatePublished: "2022-12-10 00:00:00",
				Description: models.Description{
					Value: "Christmas spice cookies are shortcrust pastry sweets flavored with vanilla, ginger and cinnamon garnished with a white chocolate ganache!",
				},
				Keywords: models.Keywords{
					Values: "recipes, recipe, italian cuisine, how to cook, Christmas spice cookies",
				},
				Image: models.Image{
					Value: "https://www.giallozafferano.com/images/260-26068/Christmas-spice-cookies_1200x800.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"Flour 00 2 &frac14; cups",
						"Butter cold ½ cup",
						"Brown sugar ½ cup",
						"Eggs 1",
						"Honey ¾ tbsp",
						"Fine salt &frac14; tbsp",
						"Vanilla extract 1 &frac14; tsp",
						"Powdered ginger 1 &frac14; tsp",
						"Cinnamon powder 1 &frac14; tsp",
						"Baking powder 1 tsp",
						"White chocolate 1 &frac14; cup",
						"Heavy cream cold ½ cup",
						"Vanilla extract 1 tbsp",
						"Powdered ginger to taste",
						"Chopped hazelnuts to taste",
						"Colored sprinkles",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"To prepare the Christmas spice cookies , start with the shortcrust pastry: in a large bowl, pour the flour, the diced butter 1 , the brown sugar 2 and the cinnamon 3 .",
						"Also add the ginger powder 4 , the salt 5 and the baking powder 6 .",
						"Start kneading with your hands 7 until you get a \"crumble\" consistency 8 . At this point add the honey 9 .",
						"Also add the egg 10 and the vanilla extract 11 . Continue kneading 12 .",
						"You will have to obtain a smooth and homogeneous mixture 13 . Transfer it on a surface and compact it to form a loaf 14 ,and place it on a sheet of parchment paper 15 .",
						"Cover with another sheet of parchment paper and immediately roll out the pastry to a thickness of about .20\" (5 mm) 16 . Transfer to the fridge and let it rest for at least 30 minutes. After that, make the cookies using a pastry ring with a diameter of 2.33\" (6 cm) 17 . Place them gradually on a baking sheet lined with parchment paper, spacing them apart 18 : with these doses you will get about 30 biscuits. Bake in a preheated static oven at 340 \u00b0 F (170\u00b0C) for about 15-20 mins.",
						"Meanwhile prepare the ganache. Coarsely chop the white chocolate 19 and melt it in the microwave or in a bain-marie. Pour the vanilla extract 20 , the ginger powder and the cold cream 21 .",
						"Mix well and transfer into a pastry bag without a nozzle 22 . Place in the refrigerator to firm up. As soon as they are ready, take the cookies out of the oven and let them cool completely 23 . Once cold you can decorate them with the ganache 24 . If the ganache is too solid, keep it at room temperature for a few minutes.",
						"Garnish with chopped hazelnuts 25 or colored sprinkles 26 . Christmas spice cookies are ready to enjoy 27 !",
					},
				},
				Name:     "Christmas spice cookies",
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 30},
				URL:      "https://www.giallozafferano.com/recipes/Christmas-spice-cookies.html",
			},
		},
		{
			name: "gimmesomeoven.com",
			in:   "https://www.gimmesomeoven.com/miso-chocolate-peanut-butter-cornflake-bars-gimme-some-oven/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DatePublished: "2022-02-18",
				Description: models.Description{
					Value: "These no-bake miso chocolate peanut butter cornflake bars are quick and easy to whip up and " +
						"ridiculously delicious. See notes above for modifications to make this recipe gluten-free and/or vegan.",
				},
				Keywords: models.Keywords{Values: ""},
				Image: models.Image{
					Value: "https://www.gimmesomeoven.com/wp-content/uploads/2022/02/Peanut-Butter-Cornflake-Bars-8-1-225x225.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 tablespoon coconut oil",
						"1 cup natural creamy peanut butter",
						"1/4 cup honey",
						"2 tablespoons white (shiro) miso paste",
						"1 teaspoon vanilla extract",
						"5 cups (5 ounces) cornflakes",
						"1 1/2 cups (9 ounces) semisweet chocolate chips",
						"flaky sea salt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Line a 9 x 9-inch baking dish with parchment paper.",
						"Heat the coconut oil in a large saucepan over medium-low heat until melted. Add peanut butter, honey, " +
							"miso paste, vanilla extract, and stir steadily for 2 minutes or until lightly warmed. Remove from " +
							"heat and stir in the cornflakes until they are evenly coated with the peanut butter sauce. Transfer the " +
							"cornflake mixture to the prepared pan and use a silicone spatula (or the flat bottom of a glass or a " +
							"measuring cup) to press the mixture down firmly and evenly until it is very compact.",
						"Heat the chocolate chips in a double boiler (or in the microwave in 10-second intervals) until completely " +
							"melted, being careful not to overcook and burn the chocolate. Immediately spread the melted chocolate " +
							"in an even layer over the cornflake mixture, using a spoon to create pretty swirls if you&#8217;d like.",
						"Refrigerate the bars for 2 to 3 hours or until firm.",
						"Serve. When you’re ready to serve the bars, carefully lift up the parchment and transfer the entire batch " +
							"to a cutting board. Use a chef’s knife to carefully cut the bars into squares. Serve immediately and enjoy!",
					},
				},
				Name:     "Miso Chocolate Peanut Butter Cornflake Bars",
				PrepTime: "PT20M",
				Yield:    models.Yield{Value: 16},
				URL:      "https://www.gimmesomeoven.com/miso-chocolate-peanut-butter-cornflake-bars-gimme-some-oven/",
			},
		},
		{
			name: "globo.com",
			in:   "https://receitas.globo.com/cheesecake-com-geleia-de-frutas-vermelhas-do-bbb-22.ghtml",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Tortas e bolos"},
				CookingMethod: models.CookingMethod{Value: "Americana"},
				Cuisine:       models.Cuisine{Value: "Americana"},
				DateModified:  "2023-12-12T07:02:11.772Z",
				DatePublished: "2022-03-30T19:43:06.164Z",
				Description: models.Description{
					Value: "Veja como fazer cheesecake com geleia de frutas vermelhas. Receita é feita em camadas, sendo a massa " +
						"feita com biscoito maisena e manteiga; recheio com leite condensado, cream cheese, ovos, sal e creme " +
						"de leite e cobertura com geleia caseira de frutas vermelhas, feita com morango, blueberry, amora, " +
						"framboesa, água e açúcar.",
				},
				Keywords: models.Keywords{
					Values: "cheesecake, lanche da tarde, recepção, aniversário",
				},
				Image: models.Image{
					Value: "https://s2-receitas.glbimg.com/XjnpBAqPQSKGTlZ6fYujfyRZ0lA=/1200x/smart/filters:cover():strip_icc()/i.s3.glbimg.com/v1/AUTH_1f540e0b94d8437dbbc39d567a1dee68/internal_photos/bs/2022/s/3/rsLexpSU6nXAgmuhfKNw/cheesecake-com-geleia-de-frutas-vermelhas-bbb22-1.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 pacote de biscoito maisena ",
						"3 colheres de sopa de margarina",
						"1 lata de leite condensado ",
						"300 gramas de cream cheese ",
						"2 ovos",
						"1 pitada de sal",
						"1 caixa de creme de leite ",
						"Meia xícara de morango",
						"Meia xícara de blueberry, também conhecido como mirtilo",
						"Meia xícara de amora",
						"Meia xícara de framboesa ",
						"1 xícara de água ",
						"1 xícara de açúcar ",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Com auxílio de um mixer, processador ou liquidificador triture os biscoitos até formar uma farofinha " +
							"bem fininha.",
						"Coloque a farofinha numa tigela e misture com a manteiga, até formar uma massa bem úmida.",
						"Forre uma forma de fundo removível com a massa.",
						"Leve ao forno com a temperatura de 180 graus Celsius por aproximadamente 10 minutos ou até ficar douradinha.",
						"Bata o leite condensado, o creme de leite, os ovos e o sal na batedeira. Reserve.",
						"Coloque todos os ingredientes na panela para cozinhar em fogo baixo, e deixe até formar uma geleia. " +
							"Deixe esfriar.",
						"Ao retirar a massa da cheesecake do forno deixe esfriar.",
						"Quando a massa estiver fria, acrescente o recheio feito previamente.",
						"Volte a torta para o forno por 25 a 30 minutos a 180 graus Celsius ou até ficar douradinha.",
						"Deixe a torta esfriar e coloque a cobertura de geleia de frutas vermelhas.",
						"Leve a cheesecake para gelar por mais ou menos duas horas.",
					},
				},
				Name:  "Cheesecake com geleia de frutas vermelhas do 'BBB 22'",
				Yield: models.Yield{Value: 4},
				URL:   "https://receitas.globo.com/cheesecake-com-geleia-de-frutas-vermelhas-do-bbb-22.ghtml",
			},
		},
		{
			name: "godt.no",
			in:   "https://www.godt.no/oppskrifter/kjoett/svin/10849/koteletter-med-paerer-i-langpanne",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "middag"},
				DateModified:  "2023-10-30T14:27:09Z",
				DatePublished: "2023-10-30T14:26:58Z",
				Description: models.Description{
					Value: "Koteletter kler mange smaker, også pærer. Brunede koteletter, pærer og sjalottløk legges sammen i en ildfast form. Supergod og enkel kosemiddag!",
				},
				Keywords: models.Keywords{Values: "Svin, Helg, Rask, Enkel"},
				Image: models.Image{
					Value: "https://cdn-yams.godt.no/api/v1/godt-recipe/images/c2/c21206f6-7995-4d90-b669-7dd33a95cd34?rule=w2000_auto",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 - 6 stykker svinekoteletter", "4 stk pære", "4 - 6 stk sjalottløk",
						"1 ts timian", "1 ss rosmarin", "olivenolje", "2 ss smør", "2 ss soyasaus",
						"salt og kvernet pepper",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Krydre kotelettene med salt og pepper og brun dem godt på begge sider i en stekepanne på middels høy varme med litt olje.",
						"Kutt pærer i grove skiver og fjern kjernehuset. Skrell og del sjalottløk i båter og fordel det i en ildfast form sammen med et par spiseskjeer smør og soyasaus. Krydre med timian, rosmarin, salt og pepper.",
						"Legg kotelettene på toppen av pærene i formen og hell over stekesjyen fra pannen.",
						"Sett formen midt i ovnen på 200 °C og stek i 10-15 minutter, avhengig av tykkelsen på kotelettene.",
						"Server med potetmos eller ris.",
					},
				},
				Name:  "Koteletter med pærer i langpanne  - Enkel kosemiddag",
				Yield: models.Yield{Value: 4},
				URL:   "https://www.godt.no/oppskrifter/kjoett/svin/10849/koteletter-med-paerer-i-langpanne",
			},
		},
		{
			name: "goodfooddiscoveries.com",
			in:   "https://goodfooddiscoveries.com/hunters-chicken/#wpzoom-premium-recipe-card",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dinner"},
				CookTime:      "PT35M",
				Cuisine:       models.Cuisine{Value: "British"},
				DatePublished: "2023-11-07T21:29:59+00:00",
				Description: models.Description{
					Value: "Hunters Chicken is a British pub classic made with chicken breasts wrapped in bacon, baked in BBQ sauce and cheese. It's a delicious and very comforting dish, that you can easily make at home.",
				},
				Keywords: models.Keywords{
					Values: "hunters chicken,chicken with bbq sauce,easy oven chicken recipes,easy chicken dinner,chicken with bacon",
				},
				Image: models.Image{
					Value: "https://goodfooddiscoveries.com/wp-content/uploads/2023/11/hunters-chicken-5.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 boneless, skinless chicken breasts",
						"8 slices streaky bacon (I prefer smoked)",
						"1 cup (240 ml) barbecue sauce (store-bought or homemade)",
						"1 cup (100 g) grated/shredded mozzarella cheese (or a mixture of mozzarella and cheddar)",
						"Fresh parsley, finely chopped (for garnish)",
						"Chips/fries or mashed potatoes, for serving",
						"Vegetables (e.g. peas, broccoli, carrots), for serving",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat your oven to 200°C (400°F).",
						"Wrap each chicken breast with 2 slices of bacon, ensuring the bacon is secure around the chicken.",
						"Place wrapped chicken breasts on a baking dish. Place in the oven and cook for 30 minutes.",
						"Remove from the oven. Drain chicken juices if needed. Brush each chicken breast generously with barbecue sauce, covering the entire breast. Sprinkle the cheese over the top of each chicken breast.",
						"Place back in the oven and bake for an additional 5 minutes until the cheese is melted and bubbly.",
						"Serve the Hunter's Chicken hot, accompanied by your favourite side dishes such as chips/fried or mashed potatoes and vegetables. Enjoy!",
					},
				},
				Name: "Hunter's Chicken (with Bacon, BBQ Sauce and Cheese)",
				NutritionSchema: models.NutritionSchema{
					Calories: "440 cal",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://goodfooddiscoveries.com/hunters-chicken/#wpzoom-premium-recipe-card",
			},
		},
		{
			name: "goodhousekeeping.com",
			in:   "https://www.goodhousekeeping.com/food-recipes/a44652479/balsamic-chicken-caprese-recipe/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "dinner"},
				CookTime:      "PT0S",
				DatePublished: "2023-08-08T15:26:30.306375Z EST",
				Description: models.Description{
					Value: "Balsamic Chicken Caprese is roasted to perfection and topped with melty mozzarella, fresh basil and ripe heirloom tomatoes.",
				},
				Keywords: models.Keywords{Values: "dinner, poultry"},
				Image: models.Image{
					Value: "https://hips.hearstapps.com/hmg-prod/images/balsamic-chicken-caprese-64c91c95d79f4.jpg?crop=1.00xw:0.668xh;0,0.192xh&resize=1200:*",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 tbsp. olive oil",
						"4 6-0z boneless, skinless chicken breasts(1 1/2 lbs total)",
						"Kosher salt and pepper", "1/4 c. balsamic vinegar",
						"6 oz. fresh mozzarella, sliced", "1/4 c. basil leaves",
						"1 vine-ripe or heirloom tomato, sliced 1/4 in. thick",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Heat oven to 400°F. Heat oil in large ovenproof skillet on medium-high. Season chicken with 1/2 teaspoon salt and 1/4 teaspoon pepper and cook until deep golden brown on 1 side, 4 to 5 minutes.",
						"Flip chicken and cook 1 minute. Reduce heat to medium, add balsamic vinegar and gently simmer until slightly thickened and syrupy, 1 to 2 minutes. Transfer skillet to oven and roast 10 minutes.",
						"Turn chicken to coat in vinegar, top with mozzarella and roast until chicken registers 165°F on instant- read thermometer and cheese begins to melt, about 2 minutes.",
						"Serve topped with basil leaves and tomato and sprinkle with salt and pepper. Drizzle with any additional balsamic glaze from skillet if desired.",
					},
				},
				Name: "Balsamic Chicken Caprese",
				NutritionSchema: models.NutritionSchema{
					Calories: "376 calories",
				},
				PrepTime: "PT0S",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.goodhousekeeping.com/food-recipes/a44652479/balsamic-chicken-caprese-recipe/",
			},
		},
		{
			name: "gonnawantseconds.com",
			in:   "https://www.gonnawantseconds.com/beef-tomato-macaroni-soup/#wprm-recipe-container-15941",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Soup"},
				CookTime:      "PT30M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2022-03-25T05:00:29+00:00",
				Description: models.Description{
					Value: "This simple but satisfying, hearty Beef and Tomato MacaroniSoup will be a repeat visitor to your dining table when the temperatures drop and appetites grow.",
				},
				Keywords: models.Keywords{
					Values: "beef soup recipes, ground beef recipes, ground beef soup recipes, macaroni soup recipes, tomato soup recipes",
				},
				Image: models.Image{
					Value: "https://www.gonnawantseconds.com/wp-content/uploads/2022/03/Beef-and-Tomato-Macaroni-Soup-01.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2 tablespoons vegetable oil",
						"1 medium yellow onion, (finely chopped)",
						"1 green bell pepper, (finely chopped)",
						"2 cloves garlic, (minced)",
						"1 pound ground beef",
						"2 teaspoons chili powder",
						"2 teaspoons dried oregano",
						"1 teaspoon salt",
						"1/2 teaspoon black pepper",
						"2 (10.75-ounce) cans condensed cream of tomato soup",
						"1 (15-ounces) can diced tomatoes (undrained)",
						"32 ounces beef broth",
						"4 cups water",
						"2 cups uncooked pasta",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Heat the vegetable oil (2 tablespoons) in a large pot over medium-high heat. Add the onions (1), green " +
							"bell pepper (1), and garlic (2 cloves) and saute until the onion mixture begins to soften " +
							"about 5-6 minutes.",
						"Add the ground beef (1 pound), crumbling with a wooden spoon, cook until there is no longer any pink. " +
							"Drain off excess fat.",
						"Add the chili powder (2 teaspoons), oregano (2 teaspoons), salt (1 teaspoon), and pepper (1/2 teaspoon) " +
							"and cook over medium heat for 1-2 minutes.",
						"Add condensed cream of tomato soup (2 (10.75-ounce) cans), diced tomatoes with their juice (1 (15-ounces) " +
							"can), beef broth (32 ounces), and water (4 cups).",
						"Bring to a boil, add pasta (2 cups). Reduce heat and cover and simmer until the pasta is just al dente. " +
							"Adjust seasoning and serve.",
					},
				},
				Name: "Beef and Tomato Macaroni\u00a0Soup",
				NutritionSchema: models.NutritionSchema{
					Calories:       "829 kcal",
					Carbohydrates:  "79 g",
					Cholesterol:    "80.5 mg",
					Fat:            "40 g",
					Fiber:          "8.5 g",
					Protein:        "36 g",
					SaturatedFat:   "16 g",
					Servings:       "1",
					Sodium:         "2643 mg",
					Sugar:          "9.5 g",
					UnsaturatedFat: "7 g",
				},
				PrepTime: "PT10M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.gonnawantseconds.com/beef-tomato-macaroni-soup/#wprm-recipe-container-15941",
			},
		},
		{
			name: "grandfrais.com",
			in:   "https://www.grandfrais.com/recettes/saute-de-lapin-sauce-chasseur",
			want: models.RecipeSchema{
				AtContext: "https://schema.org",
				AtType:    models.SchemaType{Value: "Recipe"},
				CookTime:  "PT55M",
				Description: models.Description{
					Value: "GRAND FRAIS vous propose cette délicieuse recette : Sauté de lapin sauce chasseur. Faites le plein d'idées et découvrez nos conseils et astuces pour une préparation inratable. Bon appétit !",
				},
				Image: models.Image{
					Value: "https://www.grandfrais.com/images/institBackoffice/recette/desktop/gfr-20171011171655.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 lapin d'environ 1,8 kg coupé en morceaux", "200 g de lardons",
						"250 g de champignons de Paris", "2 tomates mûres", "4 échalotes",
						"20 cl de bouillon de volaille", "30 cl de vin blanc sec", "1 bouquet garni",
						"3 gousses d'ail", "30 g de farine", "4 c. à s. d'huile d'olive",
						"Sel et poivre du moulin",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Pelez les échalotes et l'ail, émincez-les finement.\u00a0Faites revenir le tout dans l'huile dans une cocotte.",
						"Ajoutez les morceaux de lapin ainsi que les lardons.\u00a0Faites les dorer de tous les côtés puis singez. Remuez avec une cuillère en bois.",
						"Mouillez avec le vin blanc et le bouillon, ajoutez le bouquet garni, le sel, le poivre, couvrez et laissez mijoter 35 min à feu doux.",
						"Pendant ce temps, pelez et concassez grossièrement les tomates.Enlevez également le pied des champignons, lavez-les et coupez-les en lamelles.",
						"Ajoutez tomates et champignons dans la cocotte et poursuivez la cuisson pendant 20 min.",
					},
				},
				Name:     "Sauté de lapin sauce chasseur",
				PrepTime: "6 personnes",
				Yield:    models.Yield{Value: 6},
				URL:      "https://www.grandfrais.com/recettes/saute-de-lapin-sauce-chasseur",
			},
		},
		{
			name: "greatbritishchefs.com",
			in:   "https://www.greatbritishchefs.com/recipes/babecued-miso-poussin-recipe",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DatePublished: "2017-03-30T00:00:00Z",
				DateModified:  "2021-07-16T15:18:02.087Z",
				Description: models.Description{
					Value: "Scott Hallsworth's tasty barbecued poussin recipe is packed with bold Japanese flavours.",
				},
				Name: "Barbecued miso poussin with lemon, garlic and chilli dip",
				Image: models.Image{
					Value: "https://media-cdn2.greatbritishchefs.com/media/hpsovny5/img68297.whqc_1426x713q80.jpg",
				},
				Category: models.Category{Value: "Main"},
				CookTime: "PT60M",
				Ingredients: models.Ingredients{
					Values: []string{
						"2 poussin",
						"90g of brown miso paste",
						"50g of caster sugar",
						"40ml of mirin",
						"40ml of sake",
						"1 green chilli, finely chopped",
						"100ml of sake",
						"2 tbsp of dark soy sauce",
						"1 red chilli, finely chopped",
						"2 tsp Tabasco green",
						"2 tsp garlic purée, fresh",
						"2 tsp yuzu juice",
						"2 tbsp of lemon juice",
						"2 tbsp of olive oil",
						"100g of daikon radish",
						"50g of carrots",
						"50g of cucumber",
						"6 mint leaves",
						"10 coriander leaves, with a bit of stem left on",
						"10ml of yuzu juice",
						"soy sauce, to taste",
						"extra virgin olive oil, to taste",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"To begin, make the den miso for the marinade. Prepare a bain marie by filling a pan of water and sitting a " +
							"heatproof bowl snugly over the top (ensuring the bottom of the bowl doesn’t touch the water). " +
							"Whisk together the miso, sugar, mirin and sake and pour into the bowl. Cook over a high heat for about " +
							"20 minutes, stirring continuously. Remove and chill",
						"Make the marinade by mixing 100ml of the den miso and the chillies together. Use a sharp knife to cut each " +
							"poussin clean in half and make a couple of score marks, one in the fat part of the drumsticks " +
							"and one in the thighs. Marinate in the miso-chilli marinade for at least 6 hours and up to 12 hours",
						"To make the dip, whisk together all the ingredients, except the oil. Slowly whisk in the oil until " +
							"emulsified. This will keep in the fridge for a month",
						"To make the salad, thinly slice the daikon on a Japanese mandoline and layer the slices in piles of 5 or " +
							"6. Using a knife, shred very thinly. Do the same with the carrot and cucumber and mix together. " +
							"Add the mint and coriander leaves and drizzle with the yuzu, soy sauce and extra virgin olive oil",
						"Set up your barbecue and get the charcoal very hot. Once the flames start to die down a little and the " +
							"embers begin to glow, put your poussins on the grill. If you’re concerned about the poussins not " +
							"being cooked through enough and burning, take off the barbecue and finish cooking in a hot oven, about " +
							"180°C/gas mark 4 for 8–10 minutes. To check the birds are done, insert a thin metal skewer or the sharp " +
							"end of small knife into the thickest part of the thigh, pause for a couple of seconds, then hold the " +
							"skewer to your lip; if its scorching hot they're done",
						"Once cooked, serve with the daikon salad and the dip on the side",
					},
				},
				Keywords: models.Keywords{Values: "easy"},
				URL:      "https://www.greatbritishchefs.com/recipes/babecued-miso-poussin-recipe",
				Yield:    models.Yield{Value: 4},
			},
		},
		{
			name: "grimgrains.com",
			in:   "https://grimgrains.com/site/red_lentil_stew.html",
			want: models.RecipeSchema{
				AtContext: "https://schema.org",
				AtType:    models.SchemaType{Value: "Recipe"},
				Description: models.Description{
					Value: "A recipe we've prepared several times on long ocean passages, it's a one-pot recipe that is very versatile, filling and quick to prepare.How to serveWe sometimes just eat it as is, or serve it on top of pasta (as a sauce) or basmati rice. We also like to serve it with flat bread.This dish is also delicious with chili pepper flakes.Another alternative is to serve it as a side dish, to omit the lentils, add more vegetables and to cook them with the same spices in a pan. SubstitutionsThe recipe works well with most vegetables, it's easy to adapt and won't affect the taste or cooking time of the dish. Take note that adding red cabbage instead of green will alter the colour of the dish. Cooking the lentils with water instead of vegetable broth is possible, the spices add enough flavor. I also sometimes add konbu dashi as a more neutral base.For a more complex taste, add cardamom, fenugreek seeds and cinnamon.Only use shelled lentils, as whole lentils take longer to cook and the vegetables will soften too much. An alternative is to use sprouted whole brown lentils, they'll cook as quickly as red lentils and are more nutritious.",
				},
				Image: models.Image{Value: "https://grimgrains.com/media/interface/logo.png"},
				Ingredients: models.Ingredients{
					Values: []string{
						"olive oil15 ml", "yellow onion1 medium", "black pepper1.25 g",
						"cumin seeds2 g", "ground turmeric1.25 g", "carrots1 medium",
						"potatoes2, medium", "green cabbage230 g", "salt1.25 g", "red lentils100 g",
						"vegetable bouillon350 ml",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Heat a pot at medium heat, then coat bottom with 15 ml (1 tbsp) of olive oil. Add 113 g (1 small, diced) yellow onion, sautée for 5 minutes.",
						"Add black pepper, 2 g (1/2 tsp) whole cumin seeds and 1.25 g (1/4 tsp) of turmeric powder.",
						"Add 60 g (1 medium, diced) carrot, 290 g (2 medium, chopped into small cubes) potatoes and 230 g (1/4 head, chopped) of green cabbage. Mix well, lower heat, cover and cook for 5-7 minutes. Add water if vegetables are sticking to the bottom.",
						"Mix in 100 g (1/2 cup) of red lentils and 350 ml (1 1/2 cups) of vegetable bouillon. Mix well, and bring to a boil. Lower heat to a simmer, cover and cook for 10 minutes. When ready, season with salt.",
						"Eat as is, or serve over basmati rice or flat bread.",
					},
				},
				Name:     "red lentil stew",
				PrepTime: "PT20M",
				Yield:    models.Yield{Value: 2},
				URL:      "https://grimgrains.com/site/red_lentil_stew.html",
			},
		},
		{
			name: "grouprecipes.com",
			in:   "http://www.grouprecipes.com/145313/sandys-green-bean-casserole.html",
			want: models.RecipeSchema{
				AtContext: "https://schema.org",
				AtType:    models.SchemaType{Value: "Recipe"},
				CookTime:  "PT30M",
				Description: models.Description{
					Value: "My sister made this couple days ago and invited me over for supper. I thought it turned out pretty good. I got my sister's approval to share this. My sister didn't have french fried onions and use cheese since my sister was out of french fried onions. My sister has been struggling financially and it's unable to buy a lot of food since my sister's been struggling.",
				},
				Keywords: models.Keywords{
					Values: ",ground beef,onion,salt and pepper,taco seasoning mix,cream of mushroom soup,french fried onions,bake,nothing specific",
				},
				Image: models.Image{
					Value: "http://s2.grouprecipes.com/images/recipes/200/78953e17aaee3465369943d7057e7d91.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 to 1 1/2 pounds ground beef", "1 small or medium onion, diced",
						"Salt and pepper", "1 package taco seasoning",
						"1 can green beans, drain juice", "1 can cream of mushroom",
						"1 package french fried onions",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 350° F.",
						"Brown the ground beef, onion, season with salt and pepper. Drain the excess fat. Add the taco seasoning and mix well. Add the can of green beans and cream of mushroom soup. Mix well. Pour into an casserole dish. Add the french fried onions. Bake for 10 to 15 minutes",
					},
				},
				Name:  "Sandy's Green Bean Casserole Recipe",
				Yield: models.Yield{Value: 3},
				URL:   "http://www.grouprecipes.com/145313/sandys-green-bean-casserole.html",
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			test(t, tc)
		})
	}
}
