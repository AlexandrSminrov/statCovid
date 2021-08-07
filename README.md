#COVID statistics
 Библиотека собирает данные с офицалных сайтов.
 В данный момент поддерживает только РФ.

## Использовиние
##### Общая статистика 
```
total, err := statCovid.GetRuTotal()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Выявлено случаев: %v\t\t Человек выздоровело: %v\t\t Человека умерло: %v\t \nВыявлено случаев за сутки: %v\t Человек выздоровело за сутки: %v\t Человека умерло за сутки: %v\t",
		total.Sick, total.Healed, total.Died, total.SickChange, total.HealedChange, total.DiedChange)
```
##### Информация по всем регионам РФ
```
statCovid.GetRuRegions()
```

##### Поиск по региону 
```
для поиска региона Москва 

regions.SearchRuRegion("MOW") 


```

##### Буквенные коды регионов России
```
Москва                   MOW
Санкт-Петербург          SPE
Московская область       MOS
Нижегородская область    NIZ
```

##### Запрос кодов регионов России
```
regions.GetCodes()
```