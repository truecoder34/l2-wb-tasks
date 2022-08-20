Паттерн Facade относится к структурным паттернам уровня объекта.

Паттерн Facade предоставляет высокоуровневый унифицированный интерфейс в виде набора имен методов к набору взаимосвязанных классов или объектов некоторой подсистемы, что облегчает ее использование.

Разбиение сложной системы на подсистемы позволяет упростить процесс разработки, а также помогает максимально снизить зависимости одной подсистемы от другой. Однако использовать такие подсистемы становиться довольно сложно. Один из способов решения этой проблемы является паттерн Facade. Наша задача, сделать простой, единый интерфейс, через который можно было бы взаимодействовать с подсистемами.

В качестве примера можно привести интерфейс автомобиля. Современные автомобили имеют унифицированный интерфейс для водителя, под которым скрывается сложная подсистема. Благодаря применению навороченной электроники, делающей большую часть работы за водителя, тот может с лёгкостью управлять автомобилем, не задумываясь, как там все работает.

Требуется для реализации:

Класс Facade предоставляющий унифицированный доступ для классов подсистемы;
Класс подсистемы SubSystemA;
Класс подсистемы SubSystemB;
Класс подсистемы SubSystemC.
Заметьте, что фасад не является единственной точкой доступа к подсистеме, он не ограничивает возможности, которые могут понадобиться "продвинутым" пользователям, желающим работать с подсистемой напрямую.

[!] В описании паттерна применяются общие понятия, такие как Класс, Объект, Абстрактный класс. Применимо к языку Go, это Пользовательский Тип, Значение этого Типа и Интерфейс. Также в языке Go за место общепринятого наследования используется агрегирование и встраивание.

<br/>
Использование:<br/>
1) для доступа к сложной системе требуется простой интерфейс
2) система очень сложна или трудна для понимания
3) точка входа необходима для каждого уровня многоуровневого программного обеспечения
4) абстракции и реализации подсистемы тесно свяаны
<br/>
+:<br/>
1) изолирует клиентов от компонентов сложной подсистемы
<br/>
-:<br/>
1) рискует стать антипаттерном "божественный объект"