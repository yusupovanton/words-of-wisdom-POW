package repository

var quotes = [...]string{
	"People die when they are killed.",
	"Many that live deserve death. And some that die deserve life. Can you give it to them? Then do not be too eager to deal out death in judgement.",
	"It does not do to dwell on dreams and forget to live.",
	"All we have to decide is what to do with the time that is given us.",
	"We are only as strong as we are united, as weak as we are divided.",
	"In my experience, there is no such thing as luck.",
	"You only live once, but if you do it right, once is enough.",
	"A room without books is like a body without a soul.",
	"Be who you are and say what you feel, because those who mind don't matter, and those who matter don't mind.",
	"If you want to know what a man's like, take a good look at how he treats his inferiors, not his equals.",
	"The truth is often what we make of it; you heard what you wanted to hear, believed what you wanted to believe.",
	"Happiness can be found, even in the darkest of times, if one only remembers to turn on the light.",
	"So do all who live to see such times. But that is not for them to decide. All we have to decide is what to do with the time that is given us.",
	"It is better to be hated for what you are than to be loved for what you are not.",
	"No, my heart will not yet despair. We may stand, if only on one leg, or at least be left still upon our knees.",
	"If you define yourself by the power to take life, the desire to dominate, to possess… then you have nothing.",
	"Skill is the child of patience.",
	"Differences of habit and language are nothing at all if our aims are identical and our hearts are open.",
	"It matters not what someone is born, but what they grow to be.",
	"Dark times lie ahead of us and there will be a time when we must choose between what is easy and what is right.",
	"It is our choices that show what we truly are, far more than our abilities.",
	"A great leap forward often requires taking two steps back.",
	"Imperfection is beauty, madness is genius and it's better to be absolutely ridiculous than absolutely boring.",
	"The truth… It is a beautiful and terrible thing, and should therefore be treated with great caution.",
	"Do not pity the dead, pity the living. And above all, those who live without love.",
	"It’s the unknown we fear when we look upon death and darkness, nothing more.",
	"Words are, in my not-so-humble opinion, our most inexhaustible source of magic. Capable of both inflicting injury, and remedying it.",
	"Numbing the pain for a while will make it worse when you finally feel it.",
	"Life is a great sea, no matter where you turn, there is still a vast ocean of experience and opportunity awaiting you.",
	"The best of us sometimes eat our words.",
	"If you focus on what you left behind, you will never see what lies ahead!",
	"There is no monster more dangerous than a lack of compassion.",
	"Its remarkable what one can do, when one is forced to.",
	"The path that leads to what we truly desire is long and difficult, but only by following that path do we achieve our goal.",
	"Pride in one's work is an excellent quality, but it most not be carried to excess.",
	"Possess the right thinking. Only then can one receive the gifts of strength, knowledge, and peace.",
	"Some say that the path from inner turmoil begins with a friendly ear. My ear is open if you care to use it.",
	"The difference between the novice and the master is that the master has failed more times than the novice has tried.",
	"The past you've lost will never come back. I myself have made so many mistakes... But we can learn from the past so we don't repeat it.",
	"Fun times are bound to end.",
	"I study too, you know...How could I teach anyone if I didn't grow and learn day by day?",
	"As you go through life, the might current of society is bound to get in your way and there will certainly be times that things don't go as you'd hoped. When this happens, don't look to society for a cause...Do not renounce society. Frankly, you'd be wasting your time...Instead just say, 'That's life!' and muddle your way through with frustration.",
	"All of the connections we encounter in this world serve as teachers who nurture us.",
	"The most inflated egos are often the most fragile.",
	"When people get hurt, they learn to hate...When people hurt others, they become hated and racked with guilt. But knowing that pain allows people to be kind.",
	"Knowing what it feels to be in pain, is exactly why we try to be kind to others.",
	"Believe in yourselves and choose life over death. Otherwise, you've led a shameful existence.",
	"I'm telling you this because you don't get it. You think you get it, which is not same as actually getting it. Get it?",
	"This place makes me think about the mistakes I've made in the past... and I've made a lot.",
	"Even if you do succeed in getting revenge, the only thing that remains is emptiness.",
	"The next generation will always surpass the previous one. It is one of the never-ending cycles in life.",
	"To know what is right and choose to ignore it is the act of a coward.",
	"Acquiring knowledge is good, but you need experience as well.",
	"Whatever happened is in your past and doesn’t change the fact that you’re here now.",
	"If you misuse a power that’s all too great, you will only destroy yourself.",
	"When things go south, it’s ok to run away!",
	"Sometimes, we have to look beyond what we want and do what’s best.",
	"We can’t just give up just because things aren’t going the way we want them.",
	"If you don’t like your destiny, don’t accept it.",
	"Just because someone stumbles and loses their way, it doesn't mean they're lost forever. Sometimes, we all need a little help.",
	"A gift can often be a curse. Give someone wings, and they may fly too close to the sun. Give them the power of prophecy, and they may live in fear of the future. Give them the greatest gift of all, powers beyond imagination, and they may think they are meant to rule the world.",
	"If you think only with your eyes, you are easy to fool.",
	"There is one person you need to learn how to control. YOU",
	"Life will knock us down. But we choose whether or not to get back up.",
	"A strong man doesn’t need to read the future. He makes his own.",
	"Unfortunately, killing is one of those things that gets easier the more you do it.",
	"When you know about something it stops being a nightmare. When you know how to fight something, it stops being so threatening.",
	"If I'm to choose between one evil and another, then I prefer not to choose at all.",
	"Remember... magic is Chaos, Art and Science. It is a curse, a blessing and progress. It all depends on who uses magic, how they use it, and to what purpose.",
	"The thing about happiness is that you only know you had it when it's gone.",
	"Finding it though, that's not the hard part. It's letting go.",
	"Too many people have opinions on things they know nothing about. And the more ignorant they are, the more opinions they have.",
	"I wish I could stay in this moment forever. But then it wouldn’t be a moment.",
	"Society, in general, has taught for many generations that when you reach a certain age, you have to learn to stop playing.",
	"We want any effort on our part to be the winning effort. We don't want to be a drop in the bucket, we want to be the entire ocean.",
	"Yesterday is history, tomorrow is a mystery, today is a gift of God, which is why we call it the present.",
	"The man who does not read has no advantage over the man who cannot read.",
	"I may not have gone where I intended to go, but I think I have ended up where I needed to be.",
	"If you don't stand for something you will fall for anything.",
	"I am enough of an artist to draw freely upon my imagination. Imagination is more important than knowledge. Knowledge is limited. Imagination encircles the world.",
	"Any fool can be happy. It takes a man with real heart to make beauty out of the stuff that makes us weep.",
	"The world is full of magic things, patiently waiting for our senses to grow sharper.",
	"A day without sunshine is like, you know, night.",
	"Being deeply loved by someone gives you strength, while loving someone deeply gives you courage.",
	"For every minute you are angry you lose sixty seconds of happiness.",
	"It takes a great deal of bravery to stand up to our enemies, but just as much to stand up to our friends.",
	"Sometimes the questions are complicated and the answers are simple.",
	"Beauty is in the eye of the beholder and it may be necessary from time to time to give a stupid or misinformed beholder a black eye.",
	"Logic will get you from A to Z; imagination will get you everywhere.",
	"The more that you read, the more things you will know. The more that you learn, the more places you'll go.",
	"Folks are usually about as happy as they make their minds up to be.",
	"Success is not final, failure is not fatal: it is the courage to continue that counts.",
	"Life is like riding a bicycle. To keep your balance, you must keep moving.",
	"I have spent my whole life scared, frightened of things that could happen, might happen, might not happen, 50-years I spent like that. Finding myself awake at three in the morning. But you know what? Ever since my diagnosis, I sleep just fine. What I came to realise is that fear, that’s the worst of it. That’s the real enemy.",
	"Hope can be a powerful force. Maybe there's no actual magic in it, but when you know what you hope for most and hold it like a light within you, you can make things happen, almost like magic.",
	"Believe something and the Universe is on its way to being changed. Because you've changed, by believing. Once you've changed, other things start to follow. Isn't that the way it works?",
	"A well-composed book is a magic carpet on which we are wafted to a world that we cannot enter in any other way.",
	"The difference between genius and stupidity is: genius has its limits.",
	"We’re all a little weird. And life is a little weird. And when we find someone whose weirdness is compatible with ours, we join up with them and fall into mutually satisfying weirdness—and call it love—true love.",
	"Time you enjoy wasting is not wasted time.",
}
