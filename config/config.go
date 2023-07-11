// 此package放置一些全局配置变量
package config

// 常用的端口，在进行端口扫描时需要用到
var TopPorts = []int{7, 5555, 9, 13, 21, 22, 23, 25, 26, 37, 53, 79, 80, 81, 88, 106, 110, 111, 113, 119, 135, 139, 143, 144, 179, 199, 389, 427, 443, 444, 445, 465, 513, 514, 515, 543, 544, 548, 554, 587, 631, 646, 873, 888, 990, 993, 995, 1025, 1026, 1027, 1028, 1029, 1080, 1110, 1433, 1443, 1720, 1723, 1755, 1900, 2000, 2001, 2049, 2121, 2181, 2717, 3000, 3128, 3306, 3389, 3986, 4899, 5000, 5009, 5051, 5060, 5101, 5190, 5357, 5432, 5631, 5666, 5800, 5900, 6000, 6001, 6646, 7000, 7001, 7002, 7003, 7004, 7005, 7070, 8000, 8008, 8009, 8080, 8081, 8443, 8888, 9100, 9999, 10000, 11211, 32768, 49152, 49153, 49154, 49155, 49156, 49157, 8088, 9090, 8090, 8001, 82, 9080, 8082, 8089, 9000, 8002, 89, 8083, 8200, 90, 8086, 801, 8011, 8085, 9001, 9200, 8100, 8012, 85, 8084, 8070, 8091, 8003, 99, 7777, 8010, 8028, 8087, 83, 808, 38888, 8181, 800, 18080, 8099, 8899, 86, 8360, 8300, 8800, 8180, 3505, 9002, 8053, 1000, 7080, 8989, 28017, 9060, 8006, 41516, 880, 8484, 6677, 8016, 84, 7200, 9085, 5555, 8280, 1980, 8161, 9091, 7890, 8060, 6080, 8880, 8020, 889, 8881, 9081, 7007, 8004, 38501, 1010, 17, 19, 255, 1024, 1030, 1041, 1048, 1049, 1053, 1054, 1056, 1064, 1065, 1801, 2103, 2107, 2967, 3001, 3703, 5001, 5050, 6004, 8031, 10010, 10250, 10255, 6888, 87, 91, 92, 98, 1081, 1082, 1118, 1888, 2008, 2020, 2100, 2375, 3008, 6648, 6868, 7008, 7071, 7074, 7078, 7088, 7680, 7687, 7688, 8018, 8030, 8038, 8042, 8044, 8046, 8048, 8069, 8092, 8093, 8094, 8095, 8096, 8097, 8098, 8101, 8108, 8118, 8172, 8222, 8244, 8258, 8288, 8448, 8834, 8838, 8848, 8858, 8868, 8879, 8983, 9008, 9010, 9043, 9082, 9083, 9084, 9086, 9087, 9088, 9089, 9092, 9093, 9094, 9095, 9096, 9097, 9098, 9099, 9443, 9448, 9800, 9981, 9986, 9988, 9998, 10001, 10002, 10004, 10008, 12018, 12443, 14000, 16080, 18000, 18001, 18002, 18004, 18008, 18082, 18088, 18090, 18098, 19001, 20000, 20720, 21000, 21501, 21502, 28018}
