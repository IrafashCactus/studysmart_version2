package repository

import (
	"log/slog"
	"time"

	"study-smart-backend/internal/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	var userCount int64
	db.Model(&model.User{}).Count(&userCount)
	if userCount > 0 {
		slog.Info("database already has data, skipping initial seed")
		return nil
	}

	slog.Info("seeding database with Rwanda curriculum and demo data...")

	hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	hashedPassword := string(hash)

	// 1. Create School & Classrooms
	school := model.School{
		Name:       "GS Officiel de Butare",
		Code:       "GSOB-001",
		District:   "Huye",
		Sector:     "Ngoma",
		Province:   "Southern Province",
		SchoolType: "secondary",
	}
	db.Create(&school)

	classroom := model.Classroom{
		SchoolID:     school.ID,
		GradeLevel:   "S4",
		Stream:       "MPC (Math-Physics-Computer)",
		AcademicYear: 2026,
	}
	db.Create(&classroom)

	// 2. Create Users
	student := model.User{
		Email:        "student@studysmart.rw",
		PasswordHash: hashedPassword,
		FullName:     "Keza Aline",
		Role:         "student",
		LanguagePref: "en",
		AvatarURL:    "https://api.dicebear.com/7.x/bottts/svg?seed=Keza",
	}
	db.Create(&student)

	student2 := model.User{
		Email:        "manzi@studysmart.rw",
		PasswordHash: hashedPassword,
		FullName:     "Manzi Eric",
		Role:         "student",
		LanguagePref: "en",
		AvatarURL:    "https://api.dicebear.com/7.x/bottts/svg?seed=Manzi",
	}
	db.Create(&student2)

	student3 := model.User{
		Email:        "uwase@studysmart.rw",
		PasswordHash: hashedPassword,
		FullName:     "Uwase Diane",
		Role:         "student",
		LanguagePref: "rw",
		AvatarURL:    "https://api.dicebear.com/7.x/bottts/svg?seed=Diane",
	}
	db.Create(&student3)

	teacher := model.User{
		Email:        "teacher@studysmart.rw",
		PasswordHash: hashedPassword,
		FullName:     "Mugisha Jean-Paul",
		Role:         "teacher",
		LanguagePref: "en",
		AvatarURL:    "https://api.dicebear.com/7.x/personas/svg?seed=Jean",
	}
	db.Create(&teacher)

	parent := model.User{
		Email:        "parent@studysmart.rw",
		PasswordHash: hashedPassword,
		FullName:     "Gasana David",
		Role:         "parent",
		LanguagePref: "en",
		AvatarURL:    "https://api.dicebear.com/7.x/personas/svg?seed=David",
	}
	db.Create(&parent)

	admin := model.User{
		Email:        "admin@studysmart.rw",
		PasswordHash: hashedPassword,
		FullName:     "Super Admin",
		Role:         "super_admin",
		LanguagePref: "en",
		AvatarURL:    "https://api.dicebear.com/7.x/bottts/svg?seed=Admin",
	}
	db.Create(&admin)

	// 3. Gamification for Keza (Student)
	now := time.Now()
	todayStr := now.Format("2006-01-02")

	db.Create(&model.StudentXP{
		UserID:       student.ID,
		TotalXP:      840,
		CurrentLevel: 4,
		XPToNext:     220,
		UpdatedAt:    now,
	})

	db.Create(&model.Streak{
		UserID:             student.ID,
		CurrentStreak:      7,
		LongestStreak:      12,
		LastActivityDate:   todayStr,
		FreezePasses:       2,
		LastFreezeUsedDate: "",
		UpdatedAt:          now,
	})

	// Gamification for other students (for Leaderboard)
	db.Create(&model.StudentXP{UserID: student2.ID, TotalXP: 1250, CurrentLevel: 5, XPToNext: 350})
	db.Create(&model.Streak{UserID: student2.ID, CurrentStreak: 14, LongestStreak: 14, LastActivityDate: todayStr})

	db.Create(&model.StudentXP{UserID: student3.ID, TotalXP: 620, CurrentLevel: 3, XPToNext: 180})
	db.Create(&model.Streak{UserID: student3.ID, CurrentStreak: 3, LongestStreak: 5, LastActivityDate: todayStr})

	// 4. Subjects & Topics (Rwanda Curriculum)
	subMath := model.Subject{
		Name:        "Mathematics",
		NameRW:      "Imibare",
		NameFR:      "Mathématiques",
		Code:        "MATH",
		Icon:        "pi pi-calculator",
		Color:       "#3b82f6",
		GradeLevels: "S1,S2,S3,S4,S5,S6",
	}
	db.Create(&subMath)

	subPhy := model.Subject{
		Name:        "Physics",
		NameRW:      "Ubugenge",
		NameFR:      "Physique",
		Code:        "PHY",
		Icon:        "pi pi-bolt",
		Color:       "#8b5cf6",
		GradeLevels: "S1,S2,S3,S4,S5,S6",
	}
	db.Create(&subPhy)

	subBio := model.Subject{
		Name:        "Biology",
		NameRW:      "Ibinyabuzima",
		NameFR:      "Biologie",
		Code:        "BIO",
		Icon:        "pi pi-heart",
		Color:       "#10b981",
		GradeLevels: "S1,S2,S3,S4,S5,S6",
	}
	db.Create(&subBio)

	subEnt := model.Subject{
		Name:        "Entrepreneurship",
		NameRW:      "Kwihangira Imirimo",
		NameFR:      "Entrepreneuriat",
		Code:        "ENT",
		Icon:        "pi pi-briefcase",
		Color:       "#f59e0b",
		GradeLevels: "S1,S2,S3,S4,S5,S6",
	}
	db.Create(&subEnt)

	// Topics
	topicQuadratic := model.Topic{
		SubjectID:      subMath.ID,
		GradeLevel:     "S4",
		UnitNumber:     2,
		Title:          "Quadratic Equations & Discriminant",
		TitleRW:        "Inyangingo n'Imbonerahamwe ya Kabiri",
		Description:    "Roots of ax² + bx + c = 0, discriminant Δ = b² - 4ac, nature of roots.",
		CompetencyTags: "REB-MATH-S4-U02,REB-ALG-04",
		SortOrder:      1,
	}
	db.Create(&topicQuadratic)

	topicKinematics := model.Topic{
		SubjectID:      subPhy.ID,
		GradeLevel:     "S4",
		UnitNumber:     1,
		Title:          "Kinematics & Motion in 1D/2D",
		TitleRW:        "Imyirondorere y'Urugendo",
		Description:    "Velocity, acceleration, projectile trajectory, and equations of motion under gravity.",
		CompetencyTags: "REB-PHY-S4-U01,REB-MEC-01",
		SortOrder:      1,
	}
	db.Create(&topicKinematics)

	topicGenetics := model.Topic{
		SubjectID:      subBio.ID,
		GradeLevel:     "S4",
		UnitNumber:     4,
		Title:          "Mendelian Genetics & Inheritance",
		TitleRW:        "Uruhererekane rw'Irimere",
		Description:    "Monohybrid crosses, dominant vs recessive alleles, Punnett squares.",
		CompetencyTags: "REB-BIO-S4-U04",
		SortOrder:      1,
	}
	db.Create(&topicGenetics)

	topicEnt := model.Topic{
		SubjectID:      subEnt.ID,
		GradeLevel:     "S4",
		UnitNumber:     3,
		Title:          "Business Model Canvas & Market Validation",
		TitleRW:        "Igenamushinga no Gusesengura Isoko",
		Description:    "Value proposition, customer segments, break-even point in Rwandan SME context.",
		CompetencyTags: "REB-ENT-S4-U03",
		SortOrder:      1,
	}
	db.Create(&topicEnt)

	// 5. Mind-Feed Bite-Sized Cards with Inline Interactive Checks
	cards := []model.ContentCard{
		{
			TopicID:        topicQuadratic.ID,
			CardType:       "micro_concept",
			Title:          "The Discriminant (Δ): Nature of Roots",
			ContentText:    "For any quadratic equation ax² + bx + c = 0, the discriminant is Δ = b² - 4ac.\n\n• If Δ > 0: Two distinct real roots.\n• If Δ = 0: Exactly one real double root (tangent to x-axis).\n• If Δ < 0: No real roots (two complex conjugate roots).",
			KeyTakeaway:    "Remember: Δ determines whether the parabola touches, crosses, or floats above the x-axis!",
			Difficulty:     2,
			QuestionPrompt: "What is the nature of roots for 2x² - 4x + 2 = 0?",
			QuestionType:   "single_choice",
			OptionsJSON:    `["Two distinct real roots", "One repeated real root (double root)", "No real roots", "Infinitely many roots"]`,
			CorrectAnswer:  "One repeated real root (double root)",
			Explanation:    "b² - 4ac = (-4)² - 4(2)(2) = 16 - 16 = 0. Since Δ = 0, there is exactly one double root (x = 1).",
		},
		{
			TopicID:        topicKinematics.ID,
			CardType:       "micro_concept",
			Title:          "Projectile Motion: Independence of X and Y",
			ContentText:    "When an object is launched into the air in Kigali's hilly terrain, horizontal and vertical motions are completely independent:\n\n1. Horizontal (x): Constant velocity (ax = 0), vx = v₀ cos(θ)\n2. Vertical (y): Constant acceleration (ay = -g), vy = v₀ sin(θ) - gt\n\nGravity g ≈ 9.8 m/s² acts ONLY in the vertical direction.",
			KeyTakeaway:    "At maximum height, vertical velocity vy = 0 m/s, but horizontal velocity vx remains constant!",
			Difficulty:     3,
			QuestionPrompt: "A ball is thrown at 30° to the horizontal. At the very peak of its flight, what is its vertical velocity?",
			QuestionType:   "single_choice",
			OptionsJSON:    `["0 m/s", "Equal to initial velocity", "9.8 m/s", "Negative initial velocity"]`,
			CorrectAnswer:  "0 m/s",
			Explanation:    "At maximum altitude, vertical upward motion pauses instantaneously before reversing direction, so vy = 0.",
		},
		{
			TopicID:        topicGenetics.ID,
			CardType:       "flashcard",
			Title:          "Genotype vs. Phenotype",
			Front:          "What is the difference between Genotype and Phenotype?",
			Back:           "Genotype is the exact genetic makeup (alleles, e.g., Bb or BB).\n\nPhenotype is the observable physical characteristic expressed (e.g., Brown eyes, Purple flower).",
			ContentText:    "Mendel demonstrated that an organism can carry a recessive allele in its genotype without expressing it in its phenotype if paired with a dominant allele.",
			Difficulty:     1,
			QuestionPrompt: "Which of the following represents a homozygous recessive genotype?",
			QuestionType:   "single_choice",
			OptionsJSON:    `["Aa", "AA", "aa", "Ab"]`,
			CorrectAnswer:  "aa",
			Explanation:    "'Homozygous' means both alleles are identical; 'recessive' is represented by lowercase letters, so 'aa'.",
		},
		{
			TopicID:        topicEnt.ID,
			CardType:       "summary",
			Title:          "Calculating Break-Even Point (BEP)",
			ContentText:    "In Rwandan entrepreneurship, the Break-Even Point is the sales volume where Total Revenue = Total Costs (Zero Profit, Zero Loss).\n\nBEP (units) = Fixed Costs / (Selling Price per Unit - Variable Cost per Unit)\n\nContribution Margin = Selling Price - Variable Cost.",
			KeyTakeaway:    "Every unit sold after the BEP contributes directly to net profit!",
			Difficulty:     2,
			QuestionPrompt: "If monthly fixed rent is 50,000 RWF, unit price is 1,500 RWF, and variable cost is 500 RWF, how many units must be sold to break even?",
			QuestionType:   "single_choice",
			OptionsJSON:    `["25 units", "50 units", "100 units", "200 units"]`,
			CorrectAnswer:  "50 units",
			Explanation:    "BEP = 50,000 / (1,500 - 500) = 50,000 / 1,000 = 50 units.",
		},
		{
			TopicID:        topicQuadratic.ID,
			CardType:       "flashcard",
			Title:          "Quadratic Formula Flashcard",
			Front:          "State the general quadratic formula for roots of ax² + bx + c = 0.",
			Back:           "x = (-b ± √(b² - 4ac)) / (2a)",
			ContentText:    "Derived by completing the square on the standard form ax² + bx + c = 0.",
			Difficulty:     1,
			QuestionPrompt: "In the quadratic formula, the denominator is:",
			QuestionType:   "single_choice",
			OptionsJSON:    `["2a", "a", "4ac", "2b"]`,
			CorrectAnswer:  "2a",
			Explanation:    "The denominator is 2a, dividing both -b and the radical term.",
		},
	}
	for i := range cards {
		db.Create(&cards[i])
	}

	// 6. Quests (Daily Tasks & Weekly Milestones)
	quests := []model.Quest{
		{
			QuestType:       "daily",
			Title:           "Complete 2 Physics Micro-Lessons",
			Description:     "Review and finish at least 2 physics cards in your Mind-Feed stream today.",
			RequirementType: "cards_viewed",
			TargetCount:     2,
			XPReward:        40,
			SubjectCode:     "PHY",
		},
		{
			QuestType:       "daily",
			Title:           "Ace a Daily Practice Quiz",
			Description:     "Complete any topic quiz with a score of 80% or higher.",
			RequirementType: "quizzes_completed",
			TargetCount:     1,
			XPReward:        50,
			SubjectCode:     "ALL",
		},
		{
			QuestType:       "weekly",
			Title:           "Mastery Streak Champion",
			Description:     "Maintain an active study streak for 5 consecutive days this week.",
			RequirementType: "accuracy_target",
			TargetCount:     5,
			XPReward:        150,
			SubjectCode:     "ALL",
		},
		{
			QuestType:       "daily",
			Title:           "Focus Mode Sprint (25 Mins)",
			Description:     "Complete one 25-minute Pomodoro focus study block without closing the app.",
			RequirementType: "focus_minutes",
			TargetCount:     25,
			XPReward:        35,
			SubjectCode:     "ALL",
		},
	}
	for i := range quests {
		db.Create(&quests[i])
		// Link progress to Keza
		isCompleted := (i == 0) // first quest already done
		db.Create(&model.QuestProgress{
			UserID:       student.ID,
			QuestID:      quests[i].ID,
			CurrentCount: map[bool]int{true: quests[i].TargetCount, false: 0}[isCompleted],
			IsCompleted:  isCompleted,
			Claimed:      false,
		})
	}

	// 7. Practice Quizzes & Questions
	quiz1 := model.Quiz{
		Title:         "S4 Quadratic Mastery Check",
		TopicID:       &topicQuadratic.ID,
		QuizType:      "practice",
		TimeLimitSec:  300,
		QuestionCount: 3,
		IsPublished:   true,
	}
	db.Create(&quiz1)

	questions := []model.Question{
		{
			TopicID:       topicQuadratic.ID,
			QuestionType:  "mcq",
			Prompt:        "Find the roots of x² - 5x + 6 = 0.",
			OptionsJSON:   `["x = 2 and x = 3", "x = -2 and x = -3", "x = 1 and x = 6", "x = -1 and x = -6"]`,
			CorrectAnswer: "x = 2 and x = 3",
			Explanation:   "(x - 2)(x - 3) = 0 gives roots x = 2 and x = 3.",
			Difficulty:    1,
		},
		{
			TopicID:       topicQuadratic.ID,
			QuestionType:  "mcq",
			Prompt:        "If Δ < 0 for ax² + bx + c = 0, what does the graph y = ax² + bx + c look like?",
			OptionsJSON:   `["Crosses the x-axis at two points", "Does not intersect the x-axis", "Passes through the origin", "Is a straight line"]`,
			CorrectAnswer: "Does not intersect the x-axis",
			Explanation:   "When Δ < 0, there are no real solutions, meaning the curve never crosses or touches the x-axis.",
			Difficulty:    2,
		},
		{
			TopicID:       topicQuadratic.ID,
			QuestionType:  "mcq",
			Prompt:        "What is the vertex x-coordinate of the parabola y = 2x² - 8x + 5?",
			OptionsJSON:   `["x = 2", "x = -2", "x = 4", "x = 8"]`,
			CorrectAnswer: "x = 2",
			Explanation:   "Vertex x = -b / (2a) = -(-8) / (2 * 2) = 8 / 4 = 2.",
			Difficulty:    2,
		},
	}
	for i := range questions {
		db.Create(&questions[i])
	}

	// 8. Sample Mastery Scores for Keza
	db.Create(&model.MasteryScore{
		UserID:          student.ID,
		TopicID:         topicQuadratic.ID,
		MasteryPct:      88.5,
		TotalAttempts:   12,
		CorrectAttempts: 10,
		LastPracticedAt: now,
	})
	db.Create(&model.MasteryScore{
		UserID:          student.ID,
		TopicID:         topicKinematics.ID,
		MasteryPct:      64.0,
		TotalAttempts:   8,
		CorrectAttempts: 5,
		LastPracticedAt: now.Add(-24 * time.Hour),
	})
	db.Create(&model.MasteryScore{
		UserID:          student.ID,
		TopicID:         topicEnt.ID,
		MasteryPct:      92.0,
		TotalAttempts:   10,
		CorrectAttempts: 9,
		LastPracticedAt: now.Add(-48 * time.Hour),
	})

	// 9. Sample RAG Documents and Generated Content for Admin / Verification Workbench
	doc1 := model.Document{
		Filename:         "REB_Physics_S4_Curriculum_2024.pdf",
		FileURL:          "/uploads/REB_Physics_S4_Curriculum_2024.pdf",
		FileType:         "pdf",
		FileSizeBytes:    4250000,
		SubjectID:        &subPhy.ID,
		GradeLevel:       "S4",
		DocCategory:      "textbook",
		ProcessingStatus: "completed",
		ChunkCount:       48,
	}
	db.Create(&doc1)

	db.Create(&model.GeneratedContent{
		DocumentID:  doc1.ID,
		ContentType: "question",
		Title:       "Auto-Extracted MCQ: Work-Energy Theorem",
		ContentJSON: `{"prompt": "A force of 20 N moves a box across a floor by 5 meters. How much work is done?", "options": ["50 J", "100 J", "150 J", "200 J"], "correct_answer": "100 J", "explanation": "Work = Force × Displacement = 20 N × 5 m = 100 Joules."}`,
		Status:      "draft",
		ReviewNotes: "Extracted from Chapter 3, page 74 by Gemini RAG pipeline.",
	})

	db.Create(&model.GeneratedContent{
		DocumentID:  doc1.ID,
		ContentType: "flashcard",
		Title:       "Auto-Extracted Flashcard: Gravitational Potential Energy",
		ContentJSON: `{"front": "What is the formula for Gravitational Potential Energy near Earth's surface?", "back": "PE = m × g × h (Mass × Gravitational Field Strength × Height)", "keyTakeaway": "Measured in Joules (J)."}`,
		Status:      "draft",
		ReviewNotes: "Targeted for S4 Physics exam revision.",
	})

	slog.Info("database seeded successfully with full multi-role platform data!")
	return nil
}
