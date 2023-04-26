package database

import (
	"fmt"
	"time"

	"github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
)

// User seeder
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now().UTC()
	return
}

func (User) TableName() string {
	return "users"
}

type UsersSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewUsersSeeder(cfg gorm_seeder.SeederConfiguration) UsersSeeder {
	return UsersSeeder{gorm_seeder.NewSeederAbstract(cfg)}
}

func (s *UsersSeeder) Seed(db *gorm.DB) error {
	var users []User
	for i := 0; i < s.Configuration.Rows; i++ {
		indexStr := fmt.Sprint(i)
		user := User{
			Name:     "Name LastName" + indexStr,
			Email:    "foo" + indexStr + "@bar.gov",
			Password: "password-hash-string" + indexStr,
		}
		users = append(users, user)
	}
	return db.CreateInBatches(users, s.Configuration.Rows).Error
}

func (s *UsersSeeder) Clear(db *gorm.DB) error {
	entity := User{}
	return s.SeederAbstract.Delete(db, entity.TableName())
}

// Schedule seeder
func (s *Schedule) BeforeCreate(tx *gorm.DB) (err error) {
	s.CreatedAt = time.Now().UTC()
	return
}

func (Schedule) TableName() string {
	return "schedules"
}

type SchedulesSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewSchedulesSeeder(cfg gorm_seeder.SeederConfiguration) SchedulesSeeder {
	return SchedulesSeeder{gorm_seeder.NewSeederAbstract(cfg)}
}

func (s *SchedulesSeeder) Seed(db *gorm.DB) error {
	var schedules []Schedule
	for i := 0; i < s.Configuration.Rows; i++ {
		indexStr := fmt.Sprint(i)
		startIncrement := i * 7
		start := time.Now()
		start = start.AddDate(0, 0, startIncrement)
		end := start.Add(time.Minute * 60)
		schedule := Schedule{
			Title:         "Schedule title" + indexStr,
			Description:   indexStr + "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			StartDateTime: start.Format("2006-01-02 15:04:05"),
			EndDateTime:   end.Format("2006-01-02 15:04:05"),
		}
		schedules = append(schedules, schedule)
	}
	return db.CreateInBatches(schedules, s.Configuration.Rows).Error
}

func (s *SchedulesSeeder) Clear(db *gorm.DB) error {
	entity := Schedule{}
	return s.SeederAbstract.Delete(db, entity.TableName())
}

// Speaker seeder
func (u *Speaker) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now().UTC()
	return
}

func (Speaker) TableName() string {
	return "speakers"
}

type SpeakersSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewSpeakersSeeder(cfg gorm_seeder.SeederConfiguration) SpeakersSeeder {
	return SpeakersSeeder{gorm_seeder.NewSeederAbstract(cfg)}
}

func (s *SpeakersSeeder) Seed(db *gorm.DB) error {
	var speakers []Speaker
	for i := 0; i < s.Configuration.Rows; i++ {
		indexStr := fmt.Sprint(i)
		speaker := Speaker{
			Name:        "Name LastName" + indexStr,
			Topic:       "Speaker topic: " + indexStr,
			Description: indexStr + "At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi sint occaecati cupiditate non provident, similique sunt in culpa qui officia deserunt mollitia animi, id est laborum et dolorum fuga. Et harum quidem rerum facilis est et expedita distinctio.",
			ScheduleId:  i + 1,
		}
		speakers = append(speakers, speaker)
	}
	return db.CreateInBatches(speakers, s.Configuration.Rows).Error
}

func (s *SpeakersSeeder) Clear(db *gorm.DB) error {
	entity := Speaker{}
	return s.SeederAbstract.Delete(db, entity.TableName())
}

func RunSeeder() {
	usersSeeder := NewUsersSeeder(gorm_seeder.SeederConfiguration{Rows: 10})
	schedulesSeeder := NewSchedulesSeeder(gorm_seeder.SeederConfiguration{Rows: 10})
	speakersSeeder := NewSpeakersSeeder(gorm_seeder.SeederConfiguration{Rows: 10})
	seedersStack := gorm_seeder.NewSeedersStack(DB.Db)
	seedersStack.AddSeeder(&usersSeeder)
	seedersStack.AddSeeder(&schedulesSeeder)
	seedersStack.AddSeeder(&speakersSeeder)

	err := seedersStack.Seed()
	fmt.Println(err)

	// err = seedersStack.Clear()
	// fmt.Println(err)
}
