#  ATLA Api
This is a project that will give the users information about all the charaacters from the cartoon series "Avatar: The Last Airbender" (or ATLA in short).
There is 2 object in it - Episodes and Characters, and as follows, the Characters object will contain all of the main cast from the series, including the background or bit parts characters. The Episodes object will contain information about episodes and is connected with the characters' ids, so it is possible to find out which characters were in the required episode and in which episodes the character were. 

In the characters object, you will also find information about their age, gender, status(dead, deceased, etc.) and even their nation among the 4 of them - Water Tribe, Earth Kingdom, Fire Nation and Air Nomads.

In the Episodes object, you will find in which season the episode was, what were it's title, and og course, which characters appeared there. 

In this project, it will be possible to easily make some data manipulations, such as inserting, getting, modifying and deleting the character's or episode's information. 

07.03.2024 UPD: it is only possible for Characters object yet!

## Members
```
Tazhiyeva Gaukhar, 22B030591
Gaaze Verena, 22B030332

```

## API
```
POST /characters: Create new character
GET /characters/:id: Get info about character
PUT /characters/:id: Update info about character 
DELETE /characters/:id: Delete character
```

## DB Structure

```

Table characters {
  id integer [primary key]
  name text
  age integer
  status text
  nation text
  created_at timestamp
  updated_at timestamp
}

Table episodes {
  id integer
  season_id integer
  title text
  character_id integer
  created_at timestamp
  updated_at timestamp


}

Table characters_and_episodes {
  id bigserial [primary key]
  character_id bigserial
  episode_id bigserial
  created_at timestamp
  updated_at timestamp
}


Ref: characters_and_episodes.character_id < characters.id
Ref: characters_and_episodes.episode_id < episodes.id

```