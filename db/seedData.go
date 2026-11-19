package seed

type SeedOption struct {
	Name       string
	VotesCount int
}

type SeedPoll struct {
	Name    string
	Options []SeedOption
}

var SeedPolls = []SeedPoll{
	// {"Favorite programming language", []SeedOption{{"Go", 120}, {"Python", 340}, {"Rust", 95}, {"JavaScript", 410}}},
	// {"Best pizza topping", []SeedOption{{"Pepperoni", 500}, {"Mushroom", 210}, {"Pineapple", 150}, {"Olives", 80}}},
	// {"Favorite season", []SeedOption{{"Spring", 180}, {"Summer", 260}, {"Autumn", 220}, {"Winter", 140}}},
	// {"Preferred text editor", []SeedOption{{"VS Code", 600}, {"Vim", 230}, {"Neovim", 190}, {"GoLand", 170}}},
	// {"Best pet", []SeedOption{{"Dog", 700}, {"Cat", 650}, {"Fish", 90}, {"Bird", 60}}},
	// {"Favorite movie genre", []SeedOption{{"Sci-Fi", 310}, {"Comedy", 420}, {"Thriller", 280}, {"Drama", 200}}},
	// {"Coffee or tea", []SeedOption{{"Coffee", 800}, {"Tea", 450}, {"Both", 120}, {"Neither", 40}}},
	// {"Favorite sport", []SeedOption{{"Football", 520}, {"Cricket", 610}, {"Tennis", 150}, {"Basketball", 240}}},
	{"Favorite street food", []SeedOption{{"Pani puri", 540}, {"Vada pav", 410}, {"Samosa", 380}, {"Momos", 450}}},
	{"Best Indian sweet", []SeedOption{{"Gulab jamun", 520}, {"Jalebi", 340}, {"Rasgulla", 290}, {"Laddu", 260}}},
	{"Favorite coding font", []SeedOption{{"JetBrains Mono", 430}, {"Fira Code", 380}, {"Consolas", 160}, {"Source Code Pro", 140}}},
	{"Best smartwatch brand", []SeedOption{{"Apple", 410}, {"Samsung", 280}, {"Garmin", 190}, {"Fitbit", 120}}},
	{"Favorite way to stay fit", []SeedOption{{"Gym", 450}, {"Running", 330}, {"Yoga", 290}, {"Cycling", 240}}},
	{"Favorite chocolate type", []SeedOption{{"Dark", 380}, {"Milk", 520}, {"White", 90}, {"Hazelnut", 210}}},
	{"Preferred email client", []SeedOption{{"Gmail", 720}, {"Outlook", 310}, {"Apple Mail", 180}, {"Thunderbird", 60}}},
	{"Best web browser", []SeedOption{{"Chrome", 690}, {"Firefox", 240}, {"Safari", 220}, {"Brave", 180}}},
	{"Favorite anime genre", []SeedOption{{"Shonen", 480}, {"Slice of life", 230}, {"Mecha", 140}, {"Fantasy", 310}}},
	{"Favorite car type", []SeedOption{{"SUV", 520}, {"Sedan", 290}, {"Hatchback", 340}, {"Electric", 260}}},
	{"Preferred payment method", []SeedOption{{"UPI", 820}, {"Credit card", 310}, {"Cash", 120}, {"Debit card", 190}}},
	{"Favorite biryani style", []SeedOption{{"Hyderabadi", 640}, {"Lucknowi", 280}, {"Kolkata", 240}, {"Malabar", 260}}},
	{"Best search engine", []SeedOption{{"Google", 810}, {"Bing", 90}, {"DuckDuckGo", 210}, {"Perplexity", 150}}},
	{"Favorite cricket format", []SeedOption{{"T20", 540}, {"ODI", 320}, {"Test", 260}, {"T10", 70}}},
	{"Usual bedtime", []SeedOption{{"Before 10 PM", 160}, {"10 to 12", 410}, {"12 to 2 AM", 380}, {"After 2 AM", 190}}},
	{"Favorite color", []SeedOption{{"Blue", 520}, {"Black", 380}, {"Green", 290}, {"Red", 240}}},
	{"Preferred shopping platform", []SeedOption{{"Amazon", 610}, {"Flipkart", 480}, {"Myntra", 230}, {"Local stores", 170}}},
	{"Favorite pasta shape", []SeedOption{{"Penne", 340}, {"Spaghetti", 410}, {"Fusilli", 210}, {"Macaroni", 260}}},
	{"Where do you read tech news", []SeedOption{{"Hacker News", 330}, {"Reddit", 290}, {"Twitter/X", 210}, {"Newsletters", 240}}},
	{"Best first programming language", []SeedOption{{"Python", 720}, {"JavaScript", 340}, {"C", 190}, {"Java", 230}}},
	{"Preferred headphone type", []SeedOption{{"Over-ear", 410}, {"In-ear buds", 380}, {"Neckband", 150}, {"Earbuds with case", 340}}},
	{"Favorite pizza crust", []SeedOption{{"Thin", 380}, {"Thick", 220}, {"Stuffed", 290}, {"Wood-fired", 260}}},
	{"Best monsoon snack", []SeedOption{{"Pakoras", 560}, {"Corn on the cob", 340}, {"Samosa", 280}, {"Maggi", 410}}},
	{"Favorite type of tea", []SeedOption{{"Masala chai", 650}, {"Green tea", 280}, {"Iced tea", 140}, {"Black tea", 190}}},
	{"Favorite caching tool", []SeedOption{{"Redis", 720}, {"Memcached", 140}, {"In-memory map", 210}, {"Varnish", 50}}},
	{"Preferred message queue", []SeedOption{{"Kafka", 410}, {"RabbitMQ", 330}, {"NATS", 140}, {"AWS SQS", 260}}},
	{"Favorite terminal shell", []SeedOption{{"zsh", 520}, {"bash", 430}, {"fish", 190}, {"PowerShell", 110}}},
	{"Favorite Linux distro", []SeedOption{{"Ubuntu", 540}, {"Arch", 280}, {"Fedora", 190}, {"Debian", 240}}},
	{"Preferred monitor setup", []SeedOption{{"Single monitor", 280}, {"Dual monitors", 510}, {"Ultrawide", 260}, {"Laptop only", 220}}},
	{"Favorite outdoor hobby", []SeedOption{{"Trekking", 410}, {"Photography", 290}, {"Camping", 230}, {"Birdwatching", 70}}},
	{"Favorite fast food chain", []SeedOption{{"McDonald's", 380}, {"KFC", 340}, {"Subway", 210}, {"Domino's", 450}}},
	{"Favorite dessert", []SeedOption{{"Ice cream", 560}, {"Cake", 380}, {"Brownie", 340}, {"Kheer", 210}}},
	{"Favorite pancake topping", []SeedOption{{"Maple syrup", 430}, {"Honey", 290}, {"Nutella", 480}, {"Fruit", 220}}},
	{"Preferred logging approach", []SeedOption{{"Structured JSON", 460}, {"Plain text", 190}, {"Third-party service", 230}, {"print statements", 340}}},
	{"Favorite CSS approach", []SeedOption{{"Tailwind", 540}, {"Plain CSS", 290}, {"CSS Modules", 180}, {"Styled components", 140}}},
	{"Best way to spend a rainy day", []SeedOption{{"Movies", 450}, {"Reading", 280}, {"Gaming", 310}, {"Sleeping", 390}}},
	{"Favorite app category", []SeedOption{{"Productivity", 280}, {"Games", 410}, {"Social", 330}, {"Music", 360}}},
	{"Preferred Go database library", []SeedOption{{"database/sql", 420}, {"sqlx", 290}, {"GORM", 360}, {"sqlc", 240}}},
	{"Preferred spice level", []SeedOption{{"Mild", 210}, {"Medium", 480}, {"Hot", 360}, {"Extra hot", 170}}},
	{"Favorite reading format", []SeedOption{{"Paperback", 410}, {"E-book", 290}, {"Audiobook", 180}, {"PDF", 140}}},
	{"Favorite day of the week", []SeedOption{{"Friday", 410}, {"Saturday", 380}, {"Sunday", 340}, {"Monday", 20}}},
	{"Preferred deployment target", []SeedOption{{"Vercel/Netlify", 340}, {"VPS", 290}, {"Kubernetes", 260}, {"Serverless", 310}}},
	{"Most annoying bug type", []SeedOption{{"Off-by-one", 310}, {"Nil pointer", 380}, {"Race condition", 420}, {"Typo", 290}}},
	{"Favorite sandwich filling", []SeedOption{{"Paneer", 360}, {"Cheese and veggies", 410}, {"Egg", 240}, {"Chicken", 380}}},
	{"Preferred notification style", []SeedOption{{"Sound and vibrate", 190}, {"Silent badge", 380}, {"Do not disturb", 340}, {"Email digest", 230}}},
	{"Favorite weather", []SeedOption{{"Sunny", 340}, {"Rainy", 520}, {"Cold and foggy", 280}, {"Cloudy", 310}}},
	{"Best way to celebrate a birthday", []SeedOption{{"Party with friends", 410}, {"Quiet dinner", 340}, {"Trip", 280}, {"Just sleep", 150}}},
	{"Favorite Go feature", []SeedOption{{"Goroutines", 560}, {"Simple syntax", 420}, {"Fast compile", 290}, {"Interfaces", 310}}},

}