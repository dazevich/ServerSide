# ServerSide

В файле main.go идет прослушка порта 9097 и /getCourses. При обращении к нему вызывается обработчик из пакета api apiserver. Этот обработчик получает xml файл, преобразует его в JSON и отдает
