package scraper

import (
	"github.com/reaper47/recipya/internal/models"
	"testing"
)

func TestScraper_M(t *testing.T) {
	testcases := []testcase{
		{
			name: "maangchi.com",
			in:   "https://www.maangchi.com/recipe/tteokbokki",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Snack"},
				CookTime:      "P0DT0H30M0S",
				CookingMethod: models.CookingMethod{Value: "Simmering"},
				Cuisine:       models.Cuisine{Value: "Korean"},
				DateCreated:   "2007-09-11T00:56:00+00:00",
				DateModified:  "2023-07-02T14:10:36+00:00",
				DatePublished: "2007-09-11T00:56:00+00:00",
				Description: models.Description{
					Value: "Tteokbokki is chewy rice cakes cooked in a red, spicy broth. It&#039;s something I used to eat on the streets of Korea after school. My version uses anchovy stock, which combines beautifully with the spicy sauce and soft rice cakes.",
				},
				Keywords: models.Keywords{
					Values: "cylinder shaped rice cake, ddeokbokki recipe, ddukbokki, 떡볶이 조리법, korean food, Korean recipes, korean snack, Korean spicy rice cake, Korean street snack, Maangchi recipes, spicy ricecake, topokki, tteokbokki, tteokbokki recipe",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 pound of cylinder shaped rice cake, bought or homemade. (Use a little more if you’re not adding hard boiled eggs and fish cakes)",
						"4 cups of water",
						"7 large size dried anchovies, with heads and intestines removed",
						"6 x 8 inch dried kelp", "⅓ cup hot pepper paste",
						"1 tablespoon hot pepper flakes", "1 tablespoon sugar",
						"3 green onions, cut into 3 inch long pieces",
						"2 hard boiled eggs, shelled (optional)", "½ pound fish cakes (optional)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Add the water, dried anchovies, and dried kelp to a shallow pot or pan. Boil for 15 minutes over medium high heat without the lid.",
						"Combine hot pepper paste, hot pepper flakes, and sugar in a small bowl. Remove the anchovies and kelp from the pot and add the rice cake, the mixture in the bowl, the green onion, and the optional fish cakes and hard boiled eggs. The stock will be about 2\xc2\xbd cups.",
						"Stir gently with a wooden spoon when it starts to boil. Keep stirring until the rice cake turns soft and the sauce thickens and looks shiny, which should take about 10 -15 minutes. If the rice cake is not soft enough, add more water and continue stirring until soften. When you use freshly made rice cake, it takes shorter time. If you use frozen rice cake, thaw it out and soak in cold water to soften it before cooking.",
						"Remove from the heat and serve hot.",
					},
				},
				Name:  "Hot and spicy rice cake (Tteokbokki)",
				Yield: models.Yield{Value: 4},
				URL:   "https://www.maangchi.com/recipe/tteokbokki",
			},
		},
		{
			name: "madensverden.dk",
			in:   "https://madensverden.dk/durumboller-nemme-italienske-boller-med-durum-mel/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Brød"},
				CookTime:      "PT12M",
				Cuisine:       models.Cuisine{Value: "Italiensk"},
				DatePublished: "2023-04-13T06:00:58+00:00",
				Description: models.Description{
					Value: "Her er min bedste opskrift på durumboller. De klassiske italienske boller, der bages med durummel, " +
						"og som er gode til alt fra morgenbordet til en sandwich.",
				},
				Keywords: models.Keywords{Values: "boller, hjemmebagt"},
				Image: models.Image{
					Value: "https://madensverden.dk/wp-content/uploads/2019/01/durumboller-opskrift.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"50 gram gær",
						"500 gram vand",
						"500 gram durum mel",
						"150 gram manitoba hvedemel",
						"10 gram bageenzymer",
						"12 gram salt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Lunkent vand vejes af og kommes i skålen til røremaskine. Smuldr gæren i, og rør det ud med en ske.",
						"Sæt skålen fast på røremaskinen, og monter dejkrogen på maskinen.",
						"Durummel, manitoba mel, bageenzymer og salt vejes af i en skål. Bageenzymer kan som sagt udelades, og får " +
							"stadig gode og luftige durumboller.",
						"Tilsæt gradvist melblandingen i dejen, og ælt til sidst i mindst 10 minutter ved middel hastighed.",
						"Tag dejen ud på bordet, hvor den strækkes aflang og foldes sammen.",
						"Lægges tilbage i skålen, luk til med husholdningsfilm og lad dejen hæve i 15 minutter.",
						"Beklæd en bageplade med bagepapir.",
						"Tag dejen op, og den skal ikke æltes nu. Del i 12 lige store durumboller, som sættes på bagepladen. Drys " +
							"bollerne med lidt durum mel, og læg en dyb bradepande eller bageplade ovenpå.",
						"Lad durumbollerne efterhæve i 40 minutter.",
						"Forvarm ovnen til 250 grader over- og undervarme.",
						"Bag dine durumboller i cirka 12 minutter - eller indtil de er færdige.",
					},
				},
				Name: "Durumboller - nemme italienske boller med durum mel",
				NutritionSchema: models.NutritionSchema{
					Calories: "170 kcal",
					Servings: "1",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 12},
				URL:      "https://madensverden.dk/durumboller-nemme-italienske-boller-med-durum-mel/",
			},
		},
		{
			name: "madsvin.com",
			in:   "https://madsvin.com/droemmekage/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dessert"},
				Cuisine:       models.Cuisine{Value: "Dansk"},
				DatePublished: "2023-10-23T21:30:51+00:00",
				Description: models.Description{
					Value: "Svampet drømmekage med masser af fyld (tro mig, der er ikke sparet på de gode sager).",
				},
				Keywords: models.Keywords{Values: "Drømmekage"},
				Image: models.Image{
					Value: "https://madsvin.com/wp-content/uploads/2023/10/Droemmekage-fra-brovst.jpeg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"300 g hvedemel", "100 g smør (stuetempereret)", "2 tsk vaniljesukker",
						"250 g rørsukker", "4 æg", "1 knivspids salt", "3 tsk bagepulver",
						"2 dl mælk (jeg bruger letmælk)", "150 g smør", "1 dl letmælk",
						"225 g brun farin", "75 g rørsukker", "200 g kokosmel",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Tænd ovnen på over og undervarm, 200 grader.",
						"Smelt 100 g smør i en lille gryde. Hæld dernæst 2 dl mælk mælk i gryden sammen med smørret, og sæt til side.",
						"Pisk 250 g rørsukker og 4 æg luftig i en skål.",
						"I en anden skål blandes 300 g hvedemel, 3 tsk bagepulver , 1 knivspids salt og 2 tsk vaniljesukker sammen. Sigt det hele over i den luftige æggemasser, og vend det forsigtigt sammen med en dejskraber, så du ikke slår luften ud af æggemassen.",
						"Til slut hælder du smør- og mælkeblandingen i kagedejsmassen og vender det forsigtigt i dejen med en dejskraber til massen er homogen.Til sidst vendes smøren og mælken i massen og det røres igen forsigtigt sammen med en dejskraber, til dejen er ensartet.",
						"Læg et stykke bagepapir i bradepanden (tip: krøl bagepapiret sammen først, så er det nemmere at lægge ud i bradepanden bagefter).",
						"Hæld kagemassen i bradepanden og bag den midt i din ovn ved 200 grader (over- og undervarme) i 22 minutter. Mens kagen bager skal du lave fyldet.",
						"Bland 150 g smør, 1 dl letmælk, 225 g brun farin , 75 g rørsukker og 200 g kokosmel sammen i en lille gryde. Varm fyldet op ved middelvarme til massen er ensartet. Husk at røre rundt i det en gang eller to undervejs.",
						"Tag kagen ud af ovnen efter de 22 minutter, fordel fyldet ligeligt oven på kagen og sæt den tilbage i ovnen. Bag 5 minutter yderligere, tag kagen ud og lad den hvile.",
					},
				},
				Name:  "Drømmekage - Klassikeren fra Brovst",
				Yield: models.Yield{Value: 8},
				URL:   "https://madsvin.com/droemmekage/",
			},
		},
		{
			name: "marthastewart.com",
			in:   "https://www.marthastewart.com/1539828/lemon-glazed-sheet-cake",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DateModified:  "2020-11-17T14:30:47.000-05:00",
				DatePublished: "2019-05-13T17:09:07.000-04:00",
				Description: models.Description{
					Value: "This sweet, tangy Lemon-Glazed Sheet Cake has a delicate crumb and serves a crowd.",
				},
				Keywords: models.Keywords{Values: "lemon-glazed sheet cake, cake, dessert, lemon"},
				Image: models.Image{
					Value: "https://www.marthastewart.com/thmb/qGEGYT4Q4D1RG7fxw3rl448Vax8=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/lemon-glazed-sheet-cake-0519-fe9760b1-horiz-365892a83bdb4ce3bae434e35f6656fe.jpgitok28jrW8YB",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"10 tablespoons unsalted butter, room temperature, plus more for pan",
						"1.333 cups sugar",
						"1.5 teaspoons baking powder",
						"0.25 teaspoon baking soda",
						"1.25 teaspoons kosher salt",
						"2 teaspoons grated lemon zest, plus 1 tablespoon fresh juice",
						"2 large eggs, room temperature",
						"1.667 cups cake flour (not self-rising), such as Swans Down",
						"0.5 cup whole milk, room temperature",
						"0.75 cup sugar",
						"0.25 cup cornstarch",
						"0.5 teaspoon kosher salt",
						"1.5 teaspoons grated lemon zest, plus ⅔ cup fresh juice (from about 3 lemons)",
						"4 large egg yolks",
						"4 tablespoons unsalted butter, room temperature",
						"Candied lemon zest , for serving (optional)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Cake: Preheat oven to 350 degrees. Butter a 9-by-13-inch baking pan. Beat butter with sugar, baking " +
							"powder, baking soda, salt, and lemon zest on medium-high speed until light and fluffy, " +
							"about 3 minutes. Add eggs, one at a time, beating to combine after each addition and scraping " +
							"down sides as needed. Beat in lemon juice. Beat in flour in three additions, alternating with " +
							"milk and beginning and ending with flour. Scrape batter into prepared pan, smoothing top " +
							"with an offset spatula.",
						"Bake until a tester inserted in center comes out with a few moist crumbs, 30 to 35 minutes. Transfer " +
							"pan to a wire rack; let cool completely.",
						"Glaze: Combine sugar, cornstarch, salt, and lemon zest in a saucepan. Whisk in yolks, then 1 2/3 cups " +
							"water, lemon juice, and butter. Bring to a boil over mediumhigh heat, whisking constantly, and cook, " +
							"still whisking, 1 minute. Strain through a fine-mesh sieve into a heatproof bowl. Let stand 30 minutes, " +
							"whisking occasionally.",
						"Poke 20 holes in cake with a skewer; pour glaze over top. Refrigerate at least 2 hours and up to " +
							"overnight. Decorate with candied zest. (Leftovers can be refrigerated, wrapped in plastic, up to 3 days.)",
					},
				},
				Name:     "Lemon-Glazed Sheet Cake",
				PrepTime: "PT35M",
				URL:      "https://www.marthastewart.com/1539828/lemon-glazed-sheet-cake",
				Yield:    models.Yield{Value: 1},
			},
		},
		{
			name: "matprat.no",
			in:   "https://www.matprat.no/oppskrifter/tradisjon/vafler/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Dessert"},
				Cuisine:       models.Cuisine{Value: "Europa"},
				DatePublished: "01.04.2014",
				Description: models.Description{
					Value: "Vafler er alltid en suksess! Sett frem syltet&#248;y, r&#248;mme, sm&#248;r, sukker og brunost. Da  f&#229;r alle sine &#248;nsker oppfylt. Verdens beste vafler!",
				},
				Keywords: models.Keywords{
					Values: "vafler, hvetemel, melk, egg, smør, sukker, malt kardemomme, vaffel, vaffelrøre, vafler, vafler, vafler med bær, vafler, vafler, vafler, vaffeloppskrifter, oppskrift på vafler, vaffel, vaffelkake, vaffler, vaffelrøre, den store vaffeldagen",
				},
				Image: models.Image{Value: "https://images.matprat.no/uveqekyypv"},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 dl hvetemel",
						"1 dl sukker",
						"1 ts bakepulver",
						"1 ts malt kardemomme",
						"4 dl melk",
						"3 stk. egg",
						"100 g smeltet smør",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Ha alt det t&oslash;rre i en bolle og spe med litt av melken om gangen. R&oslash;r godt mellom hver gang " +
							"for &aring; f&aring; en glatt r&oslash;re uten melklumper.",
						"R&oslash;r inn eggene og tilsett smeltet sm&oslash;r. La r&oslash;ren svelle i 1/2 time. Juster r&oslash;" +
							"ren med litt vann eller melk om den er for tykk.",
						"Stek vaflene og server dem gjerne varme.",
					},
				},
				Name:  "Vafler",
				Yield: models.Yield{Value: 1},
				URL:   "https://www.matprat.no/oppskrifter/tradisjon/vafler/",
			},
		},
		{
			name: "meljoulwan.com",
			in:   "https://meljoulwan.com/2019/06/10/thai-chicken-meatballs/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Poultry"},
				DateModified:  "2019-06-05",
				DatePublished: "2019-06-10",
				Description: models.Description{
					Value: "Tender & flavorful, these meatballs have the curry flavor in a friendly shape! A secret ingredient makes a snappy outside & tender inside.",
				},
				Keywords: models.Keywords{Values: "chicken,thai"},
				Image: models.Image{
					Value: "https://meljoulwan.com/wp-content/uploads/2019/04/thai-chicken-meatballs-15-Edit-16.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 scallions", "10-12 basil leaves", "1/4 cup canned unsweetened coconut milk",
						"2 teaspoons fish sauce", "2 tablespoons green curry paste",
						"1/2 teaspoon salt", "1/2 teaspoon powdered ginger",
						"1/4 teaspoon ground black pepper", "2 pounds ground chicken",
						"juice of 1 lime", "1/4 teaspoon baking soda", "juice of 1 lime",
						"1 teaspoon coconut sugar (omit for Whole30)",
						"2 teaspoons fish sauce or coconut aminos", "1/4 teaspoon ground black pepper",
						"1/2 small head green cabbage",
						"1 carrot",
						"2 scallions",
						"1/2 cup fresh cilantro leaves",
						"small handful cashews",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Get ready. Cover a large rimmed baking sheet with aluminum foil. Preheat the oven to 400F.",
						"Prep the meatball seasonings. Mince the scallions and basil; place in a large mixing bowl. Add the coconut milk, fish sauce, curry paste, salt, ginger, and black pepper.",
						"Make the meat dough. Crumble the meat into the bowl and mix with a wooden spoon or your hands until combined. In a small bowl, mix the lime juice and baking soda; it will fizz as the baking soda reacts with the acidic lime. Pour into the bowl with the meat and mix again.",
						"Shape the meatballs. Roll about 1 tablespoon of meat into a 1-inch ball. I like to use a 1-tablespoon scoop to speed up the process. You can also opt to make the meatballs larger; they will be more tender, less crisp.",
						"Bake the meatballs. Slide into the oven and bake 15-20 minutes, until browned and sizzling.",
						"Make the cabbage salad:\u00a0In a large mixing bowl, combine the lime juice, coconut sugar, fish sauce, and black pepper; stir with a fork until the sugar is dissolved. Shred the cabbage and carrot with a food processor or box grater and slice the scallions; place them in the bowl with the dressing. Add the cilantro and toss to combine. Coarsely chop the cashews and add to the bowl.",
						"Serve the meatballs with cabbage salad on the side. Bonus points if you also make a batch of Oven-Roasted Cauliflower Rice; it can bake at the same time as the meatballs.",
					},
				},
				Name:  "Thai Chicken Meatballs w/ Cabbage Salad (Paleo, Whole30)",
				Yield: models.Yield{Value: 30},
				URL:   "https://meljoulwan.com/2019/06/10/thai-chicken-meatballs/",
			},
		},
		{
			name: "melskitchencafe.com",
			in:   "https://www.melskitchencafe.com/grilled-rosemary-ranch-chicken/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				CookTime:      "PT12M",
				DatePublished: "2021-08-23T04:00:00+00:00",
				Image: models.Image{
					Value: "https://www.melskitchencafe.com/wp-content/uploads/2021/08/rosemary-ranch-chicken6.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1/4 cup olive oil",
						"1/2 cup ranch salad dressing (see note)",
						"1/4 cup Worcestershire sauce",
						"2 tablespoons fresh lemon juice or red wine vinegar",
						"1 tablespoon finely chopped fresh rosemary",
						"1 teaspoon salt",
						"1/4 teaspoon black pepper",
						"1 1/2 - 2 pounds chicken breasts (see note)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"For the marinade, whisk together all the ingredients except the chicken until well-combined.",
						"Place chicken in shallow dish and pour the marinade over the chicken, turning the chicken to coat " +
							"evenly. Cover the dish and refrigerate for 8-12 hours.",
						"Grill the chicken on medium-high, about 3-4 minutes per side (exact time will depend on the " +
							"thickness of the chicken) until cooked through and an instant-read thermometer registers " +
							"165 degrees F at the thickest part of the chicken. Tent the chicken with foil and let rest " +
							"for 10 minutes before slicing and serving.",
					},
				},
				Name: "Rosemary Ranch Chicken",
				NutritionSchema: models.NutritionSchema{
					Calories:      "259 kcal",
					Carbohydrates: "2 g",
					Cholesterol:   "100 mg",
					Fat:           "13 g",
					Fiber:         "1 g",
					Protein:       "32 g",
					SaturatedFat:  "2 g",
					Servings:      "1",
					Sodium:        "718 mg",
					Sugar:         "1 g",
				},
				PrepTime: "PT500M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://www.melskitchencafe.com/grilled-rosemary-ranch-chicken/",
			},
		},
		{
			name: "mindmegette.hu",
			in:   "https://www.mindmegette.hu/karamellas-lavasuti.recept/",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Category:  models.Category{Value: "Aprósütemény"},
				CookTime:  "PT1H",
				Description: models.Description{
					Value: "A karamellás lávasüti egy igen csábító édesség a hét szinte bármely napján. Éppen ezért nem csak " +
						"különleges alkalmakkor kell elővenni a receptet, hanem amikor csak kedved tartja.",
				},
				Keywords: models.Keywords{
					Values: "lávasüti,lávasütemény,recept,karamella,desszert,karamellás lávasüti",
				},
				Image: models.Image{
					Value: "https://www.mindmegette.hu/images/219/O/karamelles_lava_sutik.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"0.66 bögre 1,5%-os tej",
						"0.33 bögre habtejszín",
						"0.5 bögre sózatlan vaj",
						"1 bögre kristálycukor",
						"0.5 bögre barnacukor",
						"1 tk só",
						"1 tk vaníliakivonat",
						"10 ek vaj",
						"0.75 bögre barnacukor",
						"1 db tojás",
						"1 tk vaníliakivonat",
						"2.5 bögre liszt",
						"1.5 tk sütőpor",
						"5 tk őrölt fahéj",
						"0.5 tk só",
						"fahéjas cukor",
					},
				},
				Name:     "Karamellás lávasüti receptje  |  Mindmegette.hu",
				PrepTime: "PT1H",
				Yield:    models.Yield{Value: 16},
				URL:      "https://www.mindmegette.hu/karamellas-lavasuti.recept/",
			},
		},
		{
			name: "minimalistbaker.com",
			in:   "https://minimalistbaker.com/adaptogenic-hot-chocolate-mix/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Beverage"},
				Cuisine:       models.Cuisine{Value: "Dairy-Free"},
				DatePublished: "2020-03-01T03:00:00+00:00",
				Description: models.Description{
					Value: "Low-sugar hot chocolate mix with raw cacao powder and adaptogens like reishi mushroom, maca, " +
						"ashwagandha, and he shou wu. The perfect low-caffeine cozy beverage to replace coffee or matcha.",
				},
				Keywords: models.Keywords{Values: "hot chocolate mix"},
				Image: models.Image{
					Value: "https://minimalistbaker.com/wp-content/uploads/2020/01/Adaptogenic-Hot-Chocolate-Mix-SQUARE.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"2/3 cup unsweetened cacao powder",
						"5 tsp Ashwagandha",
						"10 tsp Reishi mushroom powder",
						"10 tsp Maca",
						"2/3 cup Tocos* ((a.k.a. rice bran solubles))",
						"1 ¼ tsp ground cinnamon ((optional))",
						"2 ½ tsp He Shou Wu ((optional))",
						"Sweetener of choice ((we prefer stevia or coconut sugar to taste*))",
						"10 ounces very hot water ((or favorite dairy-free milk))",
						"3 Tbsp Adaptogenic Hot Chocolate Mix",
						"1 tsp coconut butter ((optional // if using dairy-free milk, you can omit))",
						"1 scoop collagen* ((optional // see notes for vegan option))",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"To a large jar or container add cacao powder, ashwagandha, reishi mushroom powder, maca, tocos, ground " +
							"cinnamon, and he shou wu (optional). At this point you can keep it unsweetened and sweeten to taste " +
							"per batch, or you can add sweetener. We opted for a bit of stevia (we like Trader Joe’s stevia " +
							"packets, and we added about 5, so 1/2 a packet / serving). However, not all stevia is made equal, " +
							"and commonly they’re much sweeter, so add little by little! Alternatively, you could sweeten with " +
							"coconut sugar. We’d recommend 1-2 tsp / serving, so as the recipe is written, roughly 3-7 Tbsp.",
						"Will keep stored at room temperature (preferably in a cool, dark place) up to 3 months.",
						"To a high-speed blender (or small blender), add hot water (or dairy-free milk), hot chocolate mix (3 Tbsp " +
							"per 1 serving), coconut butter (optional), and collagen (optional). Add sweetener of choice if not " +
							"added when making mix. Note: If using dairy-free milk in place of water, you can opt to skip the " +
							"tocos and coconut butter.",
						"Blend on high until creamy and frothy — about 1 minute. Taste and adjust flavor as needed, adding more " +
							"sweetener to taste, cacao for rich chocolate flavor, maca for malty flavor, or coconut butter " +
							"for coconut flavor / butteriness. The adaptogens / mushrooms can be a bit on the bitter side, " +
							"so adding more coconut butter, cacao, and sweetener will offset this.",
						"Serve and enjoy immediately. You can also make this in a big batch for the week and reheat throughout " +
							"the week as needed either on the stovetop in a saucepan, or in our go-to milk frother. Leftovers " +
							"will keep in the refrigerator up to 3-4 days (though best when fresh). Not freezer friendly.",
					},
				},
				Name: "Adaptogenic Hot Chocolate Mix",
				NutritionSchema: models.NutritionSchema{
					Calories:      "183 kcal",
					Carbohydrates: "22.7 g",
					Fat:           "8.8 g",
					Fiber:         "4 g",
					Protein:       "4 g",
					SaturatedFat:  "3.3 g",
					Servings:      "1",
					Sugar:         "3.3 g",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 10},
				URL:      "https://minimalistbaker.com/adaptogenic-hot-chocolate-mix/",
			},
		},
		{
			name: "ministryofcurry.com",
			in:   "https://ministryofcurry.com/cranberry-sauce-recipe",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Side Dish"},
				CookTime:      "PT15M",
				Cuisine:       models.Cuisine{Value: "American"},
				DatePublished: "2019-11-25T22:00:38+00:00",
				Description: models.Description{
					Value: "Hands-free 15-minute cranberry sauce recipe using only 4-ingredients is a perfect side dish for Thanksgiving feast",
				},
				Keywords: models.Keywords{Values: "cranberry sauce"},
				Image: models.Image{
					Value: "https://ministryofcurry.com/wp-content/uploads/2019/11/cranberry-sauce-2-1.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 12 oz cranberries (fresh or frozen)",
						"1 apple (peeled and diced (I like honey crisp or granny smith variety))",
						"½ cup orange juice (preferably fresh )", "¾ cup sugar", "2 twigs rosemary",
						"1 cinnamon stick", "1 tablespoon ginger (grated)",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Add cranberries, apple, and orange juice to the Instant Pot. Layer with sugar, do not mix.",
						"Pressure cook for 2 mins. Allow 5 to 10 mins natural pressure release. Open the Instant Pot. Mix well and allow to cool before serving.",
						"Add cranberries, apple, sugar, orange juice, and 1/2 cup of water to a medium saucepan. Cook on medium heat until it starts to boil and the cranberries start to pop, stirring frequently.",
						"Mash it with a potato masher or the back of the spoon to get to the texture you like. Then allow to simmer on low heat for 5 minutes.",
						"Taste the sauce and stir in more sugar, 1 tablespoon at a time to bring the sauce to your desired sweetness.",
					},
				},
				Name: "Cranberry Sauce Recipe",
				NutritionSchema: models.NutritionSchema{
					Calories:      "125 kcal",
					Carbohydrates: "32 g",
					Fat:           "1 g",
					Fiber:         "1 g",
					Protein:       "1 g",
					SaturatedFat:  "1 g",
					Servings:      "1",
					Sodium:        "1 mg",
					Sugar:         "30 g",
				},
				Yield: models.Yield{Value: 6},
				URL:   "https://ministryofcurry.com/cranberry-sauce-recipe",
			},
		},
		{
			name: "misya.info",
			in:   "https://www.misya.info/ricetta/grigliata-di-carne.htm",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Secondi di carne"},
				CookTime:      "PT25M",
				Cuisine:       models.Cuisine{Value: "Italiana"},
				DatePublished: "2022-04-03",
				Description: models.Description{
					Value: "La grigliata di carne è un grande classico del pranzo di Pasquetta. Se quest'anno avremo la " +
						"fortuna di una Pasquetta senza pioggia potreste sfruttare solo la parte iniziale del procedimento " +
						"per poi procedere con la cottura sul barbecue. In alternativa, direi che potete sfruttare la " +
						"cottura su bistecchiera come ho fatto io: non sarà proprio la stessa cosa, ma vi assicuro che " +
						"la grigliata di carne sarà comunque buonissima e ricordatevi sempre che l'importante è riuscire " +
						"a godervi la compagnia dei vostri cari ;)",
				},
				Keywords: models.Keywords{
					Values: "grigliata di carne,ricetta grigliata di carne",
				},
				Image: models.Image{
					Value: "https://www.misya.info/wp-content/uploads/2022/03/grigliata-di-carne.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 salsicce e/o spiedini",
						"4 fette di pancetta non stagionata",
						"4 braciole di maiale",
						"pomodorini",
						"patate",
						"sale",
						"pepe in grani",
						"2 spicchi di aglio",
						"1 rametto di rosmarino",
						"olio di oliva extravergine",
						"pepe rosa",
						"aglio",
						"sale",
						"olio di oliva extravergine",
						"paprica",
						"birra",
						"sale",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preparate le 3 marinate, semplicemente unendo in 3 contenitori diversi i vari ingredienti.",
						"Mettete la pancetta nella marinata al rosmarino, le braciole in quella al pepe rosa e le salsicce/spiedini " +
							"in quelle alla birra.Coprite con pellicola trasparente e lasciate riposare per 2 ore in frigorifero.",
						"Lavate le patate e bollitele per circa 25-30 minuti, finché non saranno quasi cotte, quindi scolatele, " +
							"pelatele e tagliatele a spicchi.Fate arroventare una bistecchiera sul fuoco, scolate la carne " +
							"dalla marinata e procedete con la cottura: iniziate con salsicce e spiedini che hanno una " +
							"cottura più lunga, poi aggiungete le braciole e solo alla fine la pancetta, insieme con patate " +
							"e pomodorini (ben lavati e asciugati), che insaporiranno a dovere.",
						"La grigliata di carne è pronta, servitela subito.",
					},
				},
				Name:     "Grigliata di carne",
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 4},
				URL:      "https://www.misya.info/ricetta/grigliata-di-carne.htm",
			},
		},
		{
			name: "momsdish.com",
			in:   "https://momsdish.com/khinkali",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Main Course"},
				CookTime:      "PT15M",
				Cuisine:       models.Cuisine{Value: "georgian"},
				DatePublished: "2021-03-10T03:23:33+00:00",
				Description: models.Description{
					Value: "Khinkali are super flavorful, meat-filled dumplings that are similar to soup dumplings. They reheat well, making them great for meal prep and can even be frozen! Both the dough and filling are easy to make and they’re fun to assemble.",
				},
				Keywords: models.Keywords{Values: "dumpling recipe, georgian dumplings, khinkali recipe"},
				Image: models.Image{
					Value: "https://cdn.momsdish.com/wp-content/uploads/2021/02/Khinkali-Recipe-Georgian-Dumplings-018-scaled.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"4 cups all-purpose flour",
						"2 tsp salt",
						"2 eggs",
						"1 cup water",
						"1 lb ground beef",
						"1 lb ground chicken",
						"1 medium onion (minced)",
						"1 tsp ground black pepper",
						"1 tbsp salt (adjust to taste)",
						"1 tbsp herbs ((optional))",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a large bowl combine the flour with the salt. Make a well in the center and add the eggs. Whisk together using a fork.",
						"Add the water to the center and fold the flour into the liquid. Knead the dough by hand until it feels elastic.",
						"Cover the kneaded dough and let it rest for at least 30 minutes.",
						"Combine the beef with the chicken, minced onion, salt and pepper.",
						"Roll out the dough as thin as you possibly can.",
						"Cut the dough into 3 inch circles. Place a dollop of the meat filling in the center.",
						"Pull the edge over the filling and pinch all around, forming little pockets with meat.",
						"Bring a large pot of water to a boil. Add a few of the Khinkali at a time. Once they float to the top, give them 2-4 minutes to simmer.",
						"Remove them from the water. Serve with butter, and fresh herbs.",
					},
				},
				Name: "Khinkali Recipe (Georgian Dumplings)",
				NutritionSchema: models.NutritionSchema{
					Calories:      "127 kcal",
					Carbohydrates: "13 g",
					Cholesterol:   "35 mg",
					Fat:           "5 g",
					Fiber:         "1 g",
					Protein:       "7 g",
					SaturatedFat:  "2 g",
					Servings:      "1",
					Sodium:        "412 mg",
					Sugar:         "1 g",
					TransFat:      "1 g",
				},
				PrepTime: "PT30M",
				Yield:    models.Yield{Value: 30},
				URL:      "https://momsdish.com/khinkali",
			},
		},
		{
			name: "momswithcrockpots.com",
			in:   "https://momswithcrockpots.com/crockpot-cornbread/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Bread"},
				CookTime:      "PT120M",
				DatePublished: "2018-01-01T14:50:58+00:00",
				Description: models.Description{
					Value: "Save your oven space and make this delicious Crockpot Cornbread Recipe right in your slow cooker. " +
						"Comes out perfect every time!",
				},
				Keywords: models.Keywords{Values: "holiday"},
				Image: models.Image{
					Value: "https://momswithcrockpots.com/wp-content/uploads/2018/01/Crockpot-Cornbread-1-2.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 Cup Cornmeal",
						"1 Cup Flour",
						"1 tablespoons Baking Powder",
						"1 tablespoons Sugar",
						"1/2 tsp salt",
						"1/4 cup melted butter ( or vegetable oil)",
						"1 cup Buttermilk (or milk substitute)",
						"1 egg (beaten)",
						"2 tablespoons Honey",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a medium bowl mix together all of your dry ingredients.",
						"Stir in the milk, butter/oil, egg, and honey until a batter forms.",
						"Pour your cornbread batter into a parchment paper lined 5 qt or larger crockpot.",
						"Cover and cook on high for 1 1/2 to 2 hours.",
						"Remove from crock and serve warm.",
					},
				},
				Name:     "Crockpot Cornbread",
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://momswithcrockpots.com/crockpot-cornbread/",
			},
		},
		/*{
			name: "monsieur-cuisine.com",
			in:   "https://fr.monsieur-cuisine.com/nl/recipe/concentraat-voor-runderbouillon",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Name:          "Little chocolate puddings with a molten centre",
				Yield:         models.Yield{Value: 8},
				PrepTime:      "PT15M",
				CookTime:      "PT42M",
				DatePublished: "2018-11-06",
				DateCreated:   "2018-11-06",
				Image: models.Image{
					Value: "https://www.monsieur-cuisine.com/fileadmin/_processed_/d/7/csm_24979_Rezeptfoto_925b5_dec6d60bed.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"200 g dark chocolate",
						"200 g butter",
						"6 eggs (medium)",
						"250 g sugar",
						"1 pinch of salt",
						"120 g plain flour (type 405)",
						"40 g cocoa powder",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Brush 8 ramekins with butter and keep them in the freezer for 30 minutes.",
						"Pre-heat the oven to 210 °C. Cut the chocolate into pieces and put it into the blender jug, then, with " +
							"the measuring beaker in place, chop for 5 seconds/speed setting 8. Add the butter and, with the " +
							"measuring beaker in place, melt for 4 minutes/speed setting 2/60 °C. Transfer the mixture to another " +
							"vessel and clean out the blender jug.",
						"Put the eggs, sugar and salt into the blender jug and, with the measuring beaker in place, mix for " +
							"2 minutes/speed setting 5. Then, the measuring beaker not inserted, mix for 60 seconds/speed " +
							"setting 3, pouring the chocolate and butter mixture slowly in through the filler opening. Add the " +
							"flour and cocoa powder and, with the measuring beaker in place, mix in for 30 seconds/speed setting " +
							"3. Transfer the mixture to the ramekins and bake on the second-bottom shelf of the oven for 12 minutes. " +
							"Serve the puddings immediately.",
					},
				},
				NutritionSchema: models.NutritionSchema{
					Calories:      "2366 kj / 565 kcal",
					Protein:       "9 g",
					Carbohydrates: "55 g",
					Fat:           "35 g",
				},

				URL: "https://www.monsieur-cuisine.com/en/recipes/detail/little-chocolate-puddings-with-a-molten-centre/",
			},
		},*/
		{
			name: "motherthyme.com",
			in:   "https://www.motherthyme.com/2018/06/blt-pasta-salad.html",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				DatePublished: "2018-06-29",
				Description: models.Description{
					Value: "Everything you love about a BLT tossed in this easy and delicious BLT Pasta Salad! If you like BLT's, " +
						"you're going to love this!",
				},
				Image: models.Image{
					Value: "https://www.motherthyme.com/wp-content/uploads/2018/06/BLT-PASTA-SALAD-4-225x225.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 pound bacon (cooked and crumbled)",
						"1 pound penne pasta or similar (cooked, drained and cooled)",
						"2-3 cups chopped iceberg lettuce",
						"1-2 cup halved cherry tomatoes",
						"1 cup shredded cheddar cheese",
						"2 cups croutons (optional)",
						"1 cup mayonnaise",
						"1/4 cup sour cream",
						"1/4 cup milk",
						"1 tablespoons sugar",
						"1 tablespoon grated Parmesan cheese",
						"1/2 teaspoon dried chives",
						"1/2 teaspoon garlic powder",
						"1/4 teaspoon onion powder",
						"1/2 teaspoon liquid smoke (optional but gives it a hint of smoky flavor)",
						"1/2 teaspoon salt (plus more to taste)",
						"1/4 teaspoon pepper",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"In a large bowl toss together bacon, pasta, lettuce, tomatoes and cheese.",
						"Season with a pinch of salt and pepper.",
						"In a small bowl or mason jar mix together mayonnaise, sour cream, milk, sugar, Parmesan cheese, chives, " +
							"garlic powder, onion powder, liquid smoke, salt and pepper until combined.",
						"Toss about 1/2 cup of the dressing with salad then chill salad.",
						"Before serving toss with remaining dressing, top with croutons and season with additional salt and pepper.",
					},
				},
				Name:  "BLT Pasta Salad",
				URL:   "https://www.motherthyme.com/2018/06/blt-pasta-salad.html",
				Yield: models.Yield{Value: 1},
			},
		},
		{
			name: "mybakingaddiction.com",
			in:   "https://www.mybakingaddiction.com/pistachio-pudding-cake/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Cake"},
				CookTime:      "PT45M",
				DatePublished: "2022-03-10T12:36:34+00:00",
				Description: models.Description{
					Value: "Pistachio Pudding Cake is a simple bundt cake to make any time of year. Made with a cake mix and " +
						"pistachio pudding mix, this cake can be topped with a simple glaze or any number of frostings for " +
						"a delicious crowd-pleasing dessert.",
				},
				Keywords: models.Keywords{
					Values: "bundt cake, Cake, cake mix, dessert, recipe, st patrick's day",
				},
				Image: models.Image{
					Value: "https://www.mybakingaddiction.com/wp-content/uploads/2022/03/overhead-view-sliced-pistachio-cake.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 package yellow cake mix (15.25 ounces)",
						"1 package instant pistachio pudding mix (3.4 ounces)",
						"3/4 cup sour cream",
						"3/4 cup vegetable oil",
						"3 large eggs (lightly beaten)",
						"2 teaspoons pure vanilla extract",
						"1/2 cup water",
						"½ cup roughly chopped pistachios (plus extra for garnish)",
						"1 cup powdered sugar",
						"½ teaspoon pure vanilla extract",
						"1-2 tablespoons milk",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat oven to 350°F. Spray a 10- or 12-cup bundt cake pan with baking spray. Set aside.",
						"In a large bowl, add the cake mix, pudding mix, sour cream, vegetable oil, eggs, vanilla extract, and " +
							"water. Beat on medium speed with an electric mixer for 2 minutes. Fold in the chopped pistachios.",
						"Pour the batter into the prepared cake pan. Bake for 45-50 minutes or until a toothpick inserted into " +
							"the cake comes out clean. Allow to cool in the pan for 20 minutes before turning out onto a " +
							"wire rack to cool completely.",
						"Whisk together the glaze ingredients. Drizzle over the cooled cake and top with additional chopped pistachios.",
					},
				},
				Name: "Pistachio Pudding Cake",
				NutritionSchema: models.NutritionSchema{
					Calories:       "425 kcal",
					Carbohydrates:  "55 g",
					Cholesterol:    "50 mg",
					Fat:            "21 g",
					Fiber:          "1 g",
					Protein:        "4 g",
					SaturatedFat:   "5 g",
					Servings:       "1",
					Sodium:         "454 mg",
					Sugar:          "37 g",
					TransFat:       "0.2 g",
					UnsaturatedFat: "15 g",
				},
				PrepTime: "PT5M",
				Yield:    models.Yield{Value: 12},
				URL:      "https://www.mybakingaddiction.com/pistachio-pudding-cake/",
			},
		},
		{
			name: "mykitchen101.com",
			in:   "https://mykitchen101.com/%e5%8e%9f%e5%91%b3%e7%89%9b%e6%b2%b9%e8%9b%8b%e7%b3%95/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Cake"},
				Cuisine:       models.Cuisine{Value: "baking"},
				CookTime:      "PT55M",
				Keywords:      models.Keywords{Values: "传统牛油蛋糕, 原味牛油蛋糕, 牛油蛋糕"},
				Name:          "原味牛油蛋糕 (传统牛油蛋糕)",
				DatePublished: "2017-10-19T02:26:58+00:00",
				Description: models.Description{
					Value: "牛油蛋糕是许多烘培初学者必学的蛋糕。这个原味牛油蛋糕食谱没有添加任何人造香精，所以蛋糕有着浓郁的牛油香味。食谱采用的是分蛋法来制作，所以蛋糕的组织比较细腻。",
				},
				Image: models.Image{
					Value: "https://mykitchen101.com/wp-content/uploads/2017/10/plain-butter-cake-mykitchen101-feature.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"250 g \u514b \u725b\u6cb9 (\u6709\u76d0\uff0c\u5ba4\u6e29)",
						"100 g \u514b \u7ec6\u7802\u7cd6",
						"\u00bd tsp \u8336\u5319 \u7ec6\u76d0",
						"4 \u4e2a \u86cb\u9ec4 (A\u7ea7)",
						"280 g \u514b \u4f4e\u7b4b\u9762\u7c89",
						"1 tsp \u8336\u5319 \u6ce1\u6253\u7c89 (baking powder)",
						"130 ml \u6beb\u5347 \u725b\u5976",
						"4 \u4e2a \u86cb\u767d (A\u7ea7)",
						"100 g \u514b \u7ec6\u7802\u7cd6",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"把泡打粉加入面粉拌匀后过筛备用。",
						"把烤炉预热至160°C/320°F (可在烤炉里放一盘水，约2杯，以增加湿度，可避免蛋糕表面裂开)。",
						"把蛋白放入干净的搅拌碗，以低速打至细泡。 慢慢的把砂糖加入，继续以低速打至干性发泡 (搅拌器拉起蛋白霜会形成直立状)。(温馨提示：蛋白和用来打蛋白的器具里绝对不能有任何油份和水份，建议用洗碗剂把蛋的外壳和器具清洗干净后用干净的布抹干。用一个干净的小碗来把蛋黄和蛋白分开，如果蛋黄破裂了，那个蛋白就不要用了。)",
						"把牛油、糖和盐混合，然后用中速打至乳白状。把蛋黄一个一个加入拌匀 (完全混合后才加入另一个)。",
						"分4次把面粉和牛奶轮流加入，然后继续搅拌至均匀。",
						"把⅓的蛋白霜加入面糊中，用搅拌器轻轻混合。再把其余的蛋白霜分2次加入，用刮刀轻轻拌入至完全混合。把搅拌碗敲击桌面数次以震破大气泡。 (温馨提示： 将蛋白霜和面糊混合时，动作要轻而快。橡皮刮刀从碗底往上翻转，翻拌至均匀，不可画圈搅拌，避免蛋白霜消泡。)",
						"把面糊倒入已铺上不沾烤盘纸的8吋(20-cm)方形烤盘，用刮刀把表面稍微抹平，让烤盘跌下桌面数次以震破大气泡。",
						"以160°C/320°F烘烤45分钟后, 调至180°C/355°F继续烘烤10分钟，或至完全熟透 (用木签插入蛋糕中心不粘到面糊即可)。(温馨提醒：由于每个烤炉的温度不一样，建议的时间只供参考，请依个自的烤炉调整烘烤的时间。)",
						"出炉后，让蛋糕在烤盘里冷却10分钟后才脱模。把牛油蛋糕放在铁架上至完全冷却。",
					},
				},
				NutritionSchema: models.NutritionSchema{
					Calories:       "165 kcal",
					Carbohydrates:  "17 g",
					Cholesterol:    "55 mg",
					Fat:            "10 g",
					Fiber:          "0.3 g",
					Protein:        "3 g",
					SaturatedFat:   "6 g",
					Servings:       "1",
					Sodium:         "145 mg",
					Sugar:          "9 g",
					TransFat:       "0.3 g",
					UnsaturatedFat: "4 g",
				},
				PrepTime: "PT35M",
				Yield:    models.Yield{Value: 24},
				URL:      "https://mykitchen101.com/%e5%8e%9f%e5%91%b3%e7%89%9b%e6%b2%b9%e8%9b%8b%e7%b3%95/",
			},
		},
		{
			name: "mykitchen101en.com",
			in:   "https://mykitchen101en.com/plain-butter-cake/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Cake"},
				CookTime:      "PT55M",
				Cuisine:       models.Cuisine{Value: "Baking"},
				DatePublished: "2017-10-19T07:07:37+00:00",
				Name:          "Plain Traditional Butter Cake",
				Description: models.Description{
					Value: "Butter cake is a popular recipe among baking beginners. This butter cake recipe does not have any artificial flavour added, thus the cake has a rich buttery flavour.",
				},
				Keywords: models.Keywords{Values: "butter cake, marble butter cake recipe, plain butter cake"},
				Image: models.Image{
					Value: "https://mykitchen101en.com/wp-content/uploads/2017/10/plain-butter-cake-mykitchen101en-feature.jpg",
				},
				URL: "https://mykitchen101en.com/plain-butter-cake/",
				Ingredients: models.Ingredients{
					Values: []string{
						"250 g butter ((salted, room temperature))",
						"100 g fine granulated sugar ((\u00bd cup) )",
						"\u00bd tsp salt",
						"4 egg yolks ((grade A/size: L))",
						"280 g low protein flour (cake flour) ((2 cups) )",
						"1 tsp baking powder",
						"130 ml milk ((\u00bd cup + 2 tsps) )",
						"4 egg whites ((grade A/size: L))",
						"100 g fine granulated sugar ((\u00bd cup) )",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Mix and sieve together baking powder and low protein flour.",
						"Preheat oven to 160°C/320°F (you may add a tray of water in the oven to increase the humidity, which can help preventing the cake surface from cracking).",
						"In a clean mixing bowl, whisk egg whites on low speed until tiny bubbles. Add in sugar gradually, continue whisking on low speed until stiff peak (the peaks will point straight up and hold). (Reminder: Egg whites and equipment used to beat egg whites must be grease-free and water-free. It is advisable to clean the egg shells and equipment with dishwasher detergent, then dry with a clean cloth. Use a clean small bowl to separate the egg yolk and egg white, if the egg yolk cracks, don’t use that egg white.)",
						"Combine butter, sugar and salt, then beat over medium speed until light and fluffy. Add in egg yolks gradually, beat until fully combined before adding another.",
						"Add in flour and milk alternatively in 4 batches, mix until well combined.",
						"Add ⅓ part of meringue to batter, mix gently using balloon whisk. Add the remaining meringue in 2 batches, fold in gently using rubber spatula until well mixed. Tap mixing bowl on countertop for a few times to burst large bubbles. (Reminder: When folding meringue into batter, do it quick but gentle to avoid deflating the meringue. Scrape the bottom of mixing bowl with rubber spatula and fold over, repeat folding motion until just well mixed.)",
						"Pour the batter into greased and lined 8″ (20-cm) square baking pan, smooth the surface with spatula, then drop the cake pan on countertop for a few times to burst large bubbles.",
						"Bake at 160°C/320°F for 45 minutes, then increase to 180°C/355°F and continue baking for 10 minutes, or until fully cooked (wooden stick inserted in the centre of the cake comes out clean). (Reminder: The heat for different oven is different, the suggested time is only for reference, adjust the baking time base on your oven if necessary.)",
						"Allow the butter cake to cool in the baking pan for 10 minutes before unmoulding. Let the butter cake cools completely on a wire rack.",
					},
				},
				NutritionSchema: models.NutritionSchema{
					Calories:       "165 kcal",
					Carbohydrates:  "17 g",
					Cholesterol:    "55 mg",
					Fat:            "10 g",
					Fiber:          "0.3 g",
					Protein:        "3 g",
					SaturatedFat:   "6 g",
					Servings:       "1",
					Sodium:         "145 mg",
					Sugar:          "9 g",
					TransFat:       "0.3 g",
					UnsaturatedFat: "4 g",
				},
				PrepTime: "PT35M",
				Yield:    models.Yield{Value: 24},
			},
		},
		{
			name: "myplate.gov",
			in:   "https://www.myplate.gov/recipes/supplemental-nutrition-assistance-program-snap/20-minute-chicken-creole",
			want: models.RecipeSchema{
				AtContext: atContext,
				AtType:    models.SchemaType{Value: "Recipe"},
				Name:      "20-Minute Chicken Creole",
				Image: models.Image{
					Value: "https://myplate-prod.azureedge.us/sites/default/files/styles/recipe_525_x_350_/public/2020-10/" +
						"Chicken%20Creole.jpg?itok=IX9jnbBD",
				},
				Yield:    models.Yield{Value: 8},
				CookTime: "PT20M",
				Description: models.Description{
					Value: "This Creole-inspired dish uses chili sauce and cayenne pepper to spice it up. Tomatoes, green pepper, " +
						"celery, onions and garlic spices also surround the chicken with delicious color. This main dish can " +
						"be cooked on the stovetop or with an electric skillet.",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"1 tablespoon vegetable oil",
						"2 chicken breasts (skinless, boneless)",
						"1 can diced tomatoes (14 1/2 ounces)",
						"1 cup chili sauce",
						"1 green pepper (chopped, large)",
						"2 celery stalks (chopped)",
						"1 onion (chopped)",
						"2 garlic cloves (minced)",
						"1 teaspoon dried basil",
						"1 teaspoon parsley (dried)",
						"1/4 teaspoon cayenne pepper",
						"1/4 teaspoon salt",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Wash hands with soap and water.",
						"Heat pan over medium-high heat (350 °F in an electric skillet). Add vegetable oil and chicken and " +
							"cook until the chicken reaches an internal temperature of 165 °F\u00a0(3-5 minutes).",
						"Reduce heat to medium (300 °F in electric skillet).",
						"Add tomatoes with juice, chili sauce, green pepper, celery, onion, garlic, basil, parsley, cayenne " +
							"pepper, and salt.",
						"Bring to a boil; reduce heat to low and simmer, covered for 10-15 minutes.",
						"Serve over hot, cooked rice or whole wheat pasta.",
						"Refrigerate leftovers within 2\u00a0hours.",
					},
				},
				NutritionSchema: models.NutritionSchema{
					Calories:      "77",
					Fat:           "3 g",
					SaturatedFat:  "0 g",
					Cholesterol:   "21 mg",
					Sodium:        "255 mg",
					Carbohydrates: "6 g",
					Fiber:         "2 g",
					Sugar:         "3 g",
					Protein:       "8 g",
				},
				URL: "https://www.myplate.gov/recipes/supplemental-nutrition-assistance-program-snap/20-minute-chicken-creole",
			},
		},
		{
			name: "myrecipes.com",
			in:   "https://www.myrecipes.com/recipe/quick-easy-nachos",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        models.SchemaType{Value: "Recipe"},
				Category:      models.Category{Value: "Appetizer"},
				CookTime:      "PT15M",
				Cuisine:       models.Cuisine{Value: "TexMex"},
				DateModified:  "2023-09-26T14:52:05.511-04:00",
				DatePublished: "2006-02-03T07:37:15.000-05:00",
				Description: models.Description{
					Value: "These classic Tex-Mex nachos are loaded to the max! Avoid soggy nachos by briefly baking them before topping with cheese, seasoned beef, refried beans, guacamole, and salsa. They&#39;re a great snack, party appetizer, or even casual weeknight dinner.",
				},
				Image: models.Image{
					Value: "https://www.simplyrecipes.com/thmb/_38VUZIotH7LHCImZlAMMtlBl50=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/__opt__aboutcom__coeus__resources__content_migration__simply_recipes__uploads__2019__04__Nachos-LEAD-5-ab0842bd5c3a492b989240cca869cefb.jpg",
				},
				Ingredients: models.Ingredients{
					Values: []string{
						"For the spice mix:",
						"2 tablespoons chili powder",
						"1 1/2 teaspoons kosher salt",
						"1 teaspoon granulated garlic",
						"1 teaspoon granulated onion",
						"1 teaspoon ground cumin",
						"1/2 teaspoon dried oregano",
						"1/4 teaspoon black pepper",
						"Pinch of cayenne pepper (optional)",
						"For the nachos:",
						"1 teaspoon vegetable oil",
						"1 pound ground beef (80:20 lean-to-fat ratio)",
						"16 ounces (2 cups) refried beans, canned or homemade",
						"1/4 cup water",
						"1 large bag of tortilla chips",
						"4 ounces cheddar cheese, grated (about 2 cups), plus more for topping",
						"4 ounces Colby Jack cheese, grated (about 2 cups), plus more for topping",
						"1 cup pico de gallo, store-bought or homemade , plus more for topping",
						"1/4 cup chopped cilantro",
						"1 sliced jalapeño (pickled or fresh)",
						"Optional toppings:",
						"Guacamole",
						"Salsa",
						"Sour cream",
						"Canned black olives",
						"Sliced green onions",
						"Shredded lettuce",
						"Corn",
						"Hot sauce",
					},
				},
				Instructions: models.Instructions{
					Values: []string{
						"Preheat the oven to 350°F.",
						"Make the taco spice blend: Combine all of the spices (chili powder through cayenne) together in a small bowl.",
						"Make the beef and bean topping: Heat the vegetable oil on medium high heat until it begins to shimmer. Add the ground beef to the pan and season it with all of the taco spice blend. As the meat cooks, use a spoon to break the meat up into crumbles. Cook for about 8 minutes until the meat has browned and drain the fat using a colander. Return the meat to the pan and add the refried beans and the water. Heat the mixture until the beans are smooth and warmed through. Reduce the heat to low and keep the beef-bean mixture warm while you prepare the chips.",
						"Toast the chips: On a 13x18-inch oven-safe platter or sheet pan, arrange the tortilla chips in a single layer, overlapping them slightly. Toast the chips in the preheated oven for 5 minutes, or just until you begin to smell their aroma.",
						"Assemble and bake the nachos: Carefully remove the pan from the oven and top with one half of the shredded cheeses. Allow the heat from the chips to melt the cheese slightly before topping the chips with the beef and bean mixture. Sprinkle the remaining cheese over the beef and return the pan to the oven for 5 minutes, or until the cheese has fully melted.",
						"Top and serve: Top the nachos with the pico de gallo, chopped cilantro, jalapeño slices, or any of your preferred toppings. Serve hot. Did you love the recipe? Give us some stars and leave a comment below!",
					},
				},
				Keywords: models.Keywords{
					Values: "Quick and Easy, Nachos, Refried Beans, Super Bowl, TexMex, Tortilla, Gluten-Free, Appetizer, Snack, Game Day",
				},
				Name: "The Best Nachos",
				NutritionSchema: models.NutritionSchema{
					Calories:       "1237 kcal",
					Carbohydrates:  "40 g",
					Cholesterol:    "305 mg",
					Fat:            "75 g",
					Fiber:          "7 g",
					Protein:        "98 g",
					SaturatedFat:   "29 g",
					Servings:       "Serves 6",
					Sodium:         "1432 mg",
					Sugar:          "2 g",
					UnsaturatedFat: "0 g",
				},
				PrepTime: "PT15M",
				Yield:    models.Yield{Value: 6},
				URL:      "https://www.myrecipes.com/recipe/quick-easy-nachos",
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			test(t, tc)
		})
	}
}
