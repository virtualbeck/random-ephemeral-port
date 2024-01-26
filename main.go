package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
)

var knownPorts = [2035]int{10000, 10001, 10002, 10003, 10004, 10005, 10006, 10007, 10008, 10009, 10010, 10011, 10012, 10013, 10014, 10015, 10016, 10017, 10018, 10019, 10020, 10021, 10022, 10023, 10024, 10025, 10026, 10027, 10028, 10029, 10030, 10031, 10032, 10033, 10034, 10035, 10036, 10037, 10038, 10039, 10040, 10041, 10042, 10043, 10044, 10045, 10046, 10047, 10048, 10049, 10050, 10051, 10052, 10053, 10054, 10055, 10056, 10057, 10058, 10059, 10060, 10061, 10062, 10063, 10064, 10065, 10066, 10067, 10068, 10069, 10070, 10071, 10072, 10073, 10074, 10075, 10076, 10077, 10078, 10079, 10080, 10081, 10082, 10083, 10084, 10085, 10086, 10087, 10088, 10089, 10090, 10091, 10092, 10093, 10094, 10095, 10096, 10097, 10098, 10099, 10100, 10101, 10102, 10103, 10104, 10105, 10106, 10107, 10108, 10109, 10110, 10111, 10112, 10113, 10114, 10115, 10116, 10117, 10118, 10119, 10120, 10121, 10122, 10123, 10124, 10125, 10126, 10127, 10128, 10129, 10130, 10131, 10132, 10133, 10134, 10135, 10136, 10137, 10138, 10139, 10140, 10141, 10142, 10143, 10144, 10145, 10146, 10147, 10148, 10149, 10150, 10151, 10152, 10153, 10154, 10155, 10156, 10157, 10158, 10159, 10160, 10161, 10162, 10163, 10164, 10165, 10166, 10167, 10168, 10169, 10170, 10171, 10172, 10173, 10174, 10175, 10176, 10177, 10178, 10179, 10180, 10181, 10182, 10183, 10184, 10185, 10186, 10187, 10188, 10189, 10190, 10191, 10192, 10193, 10194, 10195, 10196, 10197, 10198, 10199, 10200, 10201, 10202, 10203, 10204, 10212, 1024, 1025, 1027, 1028, 1029, 10308, 10346, 10468, 10480, 10505, 10514, 10578, 1058, 1059, 1080, 10800, 10823, 10836, 1085, 10891, 10933, 1098, 1099, 11001, 1109, 11100, 11111, 11112, 1112, 1113, 1119, 11211, 11214, 11215, 11235, 11311, 11371, 1167, 11753, 1194, 1198, 12000, 12012, 12013, 12035, 12043, 12046, 1212, 1214, 1220, 12201, 12222, 12223, 12307, 12308, 1234, 12345, 1241, 12443, 12489, 1270, 1293, 12975, 13000, 13001, 13002, 13003, 13004, 13005, 13006, 13007, 13008, 13009, 13010, 13011, 13012, 13013, 13014, 13015, 13016, 13017, 13018, 13019, 13020, 13021, 13022, 13023, 13024, 13025, 13026, 13027, 13028, 13029, 13030, 13031, 13032, 13033, 13034, 13035, 13036, 13037, 13038, 13039, 13040, 13041, 13042, 13043, 13044, 13045, 13046, 13047, 13048, 13049, 13050, 13075, 1311, 1314, 1319, 1337, 13400, 1341, 1344, 1352, 1360, 13698, 13720, 13721, 13724, 13782, 13783, 13785, 13786, 1414, 1417, 1418, 1419, 1420, 1431, 1433, 1434, 14550, 14567, 14652, 1476, 14800, 1481, 1492, 1494, 1500, 15000, 15009, 1501, 15010, 1503, 1512, 1513, 1521, 1524, 1527, 1533, 1534, 15345, 1540, 1541, 1542, 15441, 1545, 1547, 1550, 15567, 1560, 1561, 1562, 1563, 1564, 1565, 1566, 1567, 15672, 1568, 1569, 1570, 1571, 1572, 1573, 1574, 1575, 1576, 1577, 1578, 1579, 1580, 1581, 1582, 1583, 1584, 1585, 1586, 1587, 1588, 1589, 1590, 16000, 1604, 16080, 16200, 16225, 16250, 1626, 16261, 1627, 1628, 1629, 16300, 16384, 16385, 16386, 16387, 16393, 16394, 16395, 16396, 16397, 16398, 16399, 16400, 16401, 16402, 16403, 16404, 16405, 16406, 16407, 16408, 16409, 16410, 16411, 16412, 16413, 16414, 16415, 16416, 16417, 16418, 16419, 16420, 16421, 16422, 16423, 16424, 16425, 16426, 16427, 16428, 16429, 16430, 16431, 16432, 16433, 16434, 16435, 16436, 16437, 16438, 16439, 16440, 16441, 16442, 16443, 16444, 16445, 16446, 16447, 16448, 16449, 1645, 16450, 16451, 16452, 16453, 16454, 16455, 16456, 16457, 16458, 16459, 1646, 16460, 16461, 16462, 16463, 16464, 16465, 16466, 16467, 16468, 16469, 16470, 16471, 16472, 16567, 1666, 16666, 16677, 1677, 1688, 17000, 1701, 17011, 1707, 1714, 1715, 1716, 1717, 1718, 1719, 1720, 1721, 1722, 17224, 17225, 1723, 1724, 1725, 1726, 1727, 1728, 1729, 1730, 1731, 1732, 1733, 17333, 1734, 1735, 1736, 1737, 1738, 1739, 1740, 1741, 1742, 1743, 1744, 1745, 1746, 1747, 17472, 17474, 17475, 1748, 1749, 1750, 17500, 1751, 1752, 1753, 1754, 1755, 1756, 1757, 1758, 1759, 1760, 1761, 1762, 1763, 1764, 1776, 17777, 1783, 1801, 18080, 18081, 18091, 18092, 18104, 1812, 1813, 18200, 18201, 18206, 18300, 18301, 18306, 18333, 18400, 18401, 18505, 18506, 18605, 18606, 1863, 18676, 1880, 1883, 1900, 19000, 19001, 19132, 19133, 19150, 19226, 19294, 19295, 19302, 1935, 19531, 19532, 1965, 1967, 1972, 19788, 19812, 19813, 19814, 1984, 1985, 1998, 19999, 2000, 20000, 2010, 2033, 2049, 2056, 20560, 20582, 20583, 20595, 2080, 20808, 2082, 2083, 2086, 2087, 2095, 2096, 2100, 2101, 2102, 21025, 2103, 2104, 21064, 2123, 2142, 2152, 2159, 2181, 2195, 2196, 2197, 22000, 2210, 2211, 22136, 2221, 2222, 22222, 2223, 2224, 2225, 2226, 22347, 22350, 22351, 2240, 2261, 2262, 2302, 2303, 2305, 23073, 23399, 2351, 23513, 2368, 2369, 2370, 2372, 2375, 2376, 2377, 2379, 2380, 2389, 2399, 2401, 2404, 2424, 2427, 24441, 24444, 24465, 2447, 24554, 2456, 2459, 2480, 24800, 2483, 2484, 24842, 2500, 2501, 2535, 2541, 2546, 2547, 2548, 25565, 25575, 25600, 25734, 25826, 2593, 2598, 2599, 26000, 2628, 2638, 26822, 26900, 26901, 26909, 26910, 26911, 27000, 27001, 27002, 27003, 27004, 27005, 27006, 27007, 27008, 27009, 27015, 27016, 27017, 27018, 27019, 27020, 27021, 27022, 27023, 27024, 27025, 27026, 27027, 27028, 27029, 27030, 27031, 27032, 27033, 27034, 27035, 27036, 2710, 2727, 27374, 27500, 27501, 27502, 27503, 27504, 27505, 27506, 27507, 27508, 27509, 27510, 27511, 27512, 27513, 27514, 27515, 27516, 27517, 27518, 27519, 27520, 27521, 27522, 27523, 27524, 27525, 27526, 27527, 27528, 27529, 27530, 27531, 27532, 27533, 27534, 27535, 27536, 27537, 27538, 27539, 27540, 27541, 27542, 27543, 27544, 27545, 27546, 27547, 27548, 27549, 27550, 27551, 27552, 27553, 27554, 27555, 27556, 27557, 27558, 27559, 27560, 27561, 27562, 27563, 27564, 27565, 27566, 27567, 27568, 27569, 27570, 27571, 27572, 27573, 27574, 27575, 27576, 27577, 27578, 27579, 27580, 27581, 27582, 27583, 27584, 27585, 27586, 27587, 27588, 27589, 2759, 27590, 27591, 27592, 27593, 27594, 27595, 27596, 27597, 27598, 27599, 27600, 27601, 27602, 27603, 27604, 27605, 27606, 27607, 27608, 27609, 27610, 27611, 27612, 27613, 27614, 27615, 27616, 27617, 27618, 27619, 27620, 27621, 27622, 27623, 27624, 27625, 27626, 27627, 27628, 27629, 27630, 27631, 27632, 27633, 27634, 27635, 27636, 27637, 27638, 27639, 27640, 27641, 27642, 27643, 27644, 27645, 27646, 27647, 27648, 27649, 27650, 27651, 27652, 27653, 27654, 27655, 27656, 27657, 27658, 27659, 27660, 27661, 27662, 27663, 27664, 27665, 27666, 27667, 27668, 27669, 27670, 27671, 27672, 27673, 27674, 27675, 27676, 27677, 27678, 27679, 27680, 27681, 27682, 27683, 27684, 27685, 27686, 27687, 27688, 27689, 27690, 27691, 27692, 27693, 27694, 27695, 27696, 27697, 27698, 27699, 27700, 27701, 27702, 27703, 27704, 27705, 27706, 27707, 27708, 27709, 27710, 27711, 27712, 27713, 27714, 27715, 27716, 27717, 27718, 27719, 27720, 27721, 27722, 27723, 27724, 27725, 27726, 27727, 27728, 27729, 27730, 27731, 27732, 27733, 27734, 27735, 27736, 27737, 27738, 27739, 27740, 27741, 27742, 27743, 27744, 27745, 27746, 27747, 27748, 27749, 2775, 27750, 27751, 27752, 27753, 27754, 27755, 27756, 27757, 27758, 27759, 27760, 27761, 27762, 27763, 27764, 27765, 27766, 27767, 27768, 27769, 27770, 27771, 27772, 27773, 27774, 27775, 27776, 27777, 27778, 27779, 27780, 27781, 27782, 27783, 27784, 27785, 27786, 27787, 27788, 27789, 27790, 27791, 27792, 27793, 27794, 27795, 27796, 27797, 27798, 27799, 27800, 27801, 27802, 27803, 27804, 27805, 27806, 27807, 27808, 27809, 27810, 27811, 27812, 27813, 27814, 27815, 27816, 27817, 27818, 27819, 27820, 27821, 27822, 27823, 27888, 27901, 27902, 27903, 27904, 27905, 27906, 27907, 27908, 27909, 27910, 27950, 27960, 27961, 27962, 27963, 27964, 27965, 27966, 27967, 27968, 27969, 28000, 28001, 28015, 28016, 2809, 2811, 28200, 28260, 2827, 28443, 28769, 28770, 28771, 28785, 28786, 28852, 28910, 28960, 29000, 29070, 2944, 2945, 2947, 2948, 2949, 2967, 2989, 29900, 29901, 29920, 3000, 30000, 30003, 30004, 3001, 30033, 3004, 3010, 30120, 3020, 3050, 3052, 30564, 3074, 3101, 3128, 31337, 31416, 31438, 31457, 32137, 3225, 3233, 32400, 3260, 3268, 3269, 32764, 3283, 32887, 3290, 32976, 3305, 3306, 3323, 3332, 3333, 33434, 3344, 3351, 33848, 3386, 3389, 3396, 34000, 3412, 34197, 3423, 3424, 3435, 3455, 3478, 3479, 3480, 3483, 3493, 3503, 3516, 3527, 3535, 35357, 3544, 3551, 3601, 3632, 36330, 3645, 3655, 3659, 3667, 3671, 3689, 3690, 37008, 3702, 3721, 3724, 3725, 3749, 3768, 3784, 3785, 3799, 3804, 3825, 3826, 3830, 3835, 3856, 3868, 3872, 3880, 3900, 3960, 3962, 3978, 3979, 3999, 4000, 40000, 4001, 4018, 4035, 4045, 4050, 4069, 4070, 4089, 4090, 4093, 4096, 4105, 4111, 41121, 4116, 4123, 41230, 4125, 4172, 41794, 41795, 41796, 41797, 4190, 4195, 4197, 4198, 4200, 4201, 42081, 4222, 4226, 4242, 4243, 4244, 42590, 42999, 4303, 4307, 43110, 4321, 43594, 43595, 44123, 44405, 4444, 4445, 44818, 4486, 4488, 4500, 4502, 4503, 4504, 4505, 4506, 4534, 4560, 4567, 4569, 4604, 4605, 4610, 4611, 4612, 4613, 4614, 4615, 4616, 4617, 4618, 4619, 4620, 4621, 4622, 4623, 4624, 4625, 4626, 4627, 4628, 4629, 4630, 4631, 4632, 4633, 4634, 4635, 4636, 4637, 4638, 4639, 4640, 4662, 4664, 4672, 4711, 4713, 4723, 4724, 4728, 4730, 4739, 4747, 4753, 4757, 47808, 47809, 47810, 47811, 47812, 47813, 47814, 47815, 47816, 47817, 47818, 47819, 47820, 47821, 47822, 47823, 4789, 4791, 4840, 4843, 4847, 4848, 48556, 48656, 48657, 4894, 49151, 4944, 4949, 4950, 5000, 5001, 5002, 5003, 5004, 5005, 5006, 5007, 5008, 5009, 5010, 5011, 5012, 5013, 5014, 5015, 5016, 5017, 5018, 5019, 5020, 5021, 5022, 5023, 5024, 5025, 5026, 5027, 5028, 5029, 5030, 5031, 5032, 5033, 5034, 5035, 5036, 5037, 5038, 5039, 5040, 5041, 5042, 5043, 5044, 5045, 5046, 5047, 5048, 5049, 5050, 5051, 5052, 5053, 5054, 5055, 5056, 5057, 5058, 5059, 5060, 5061, 5062, 5063, 5064, 5065, 5070, 5080, 5084, 5085, 5090, 5093, 5099, 5104, 5121, 5124, 5125, 5150, 5151, 5154, 5172, 5173, 5190, 5198, 5199, 5200, 5201, 5222, 5223, 5228, 5231, 5232, 5235, 5242, 5243, 5246, 5247, 5269, 5280, 5281, 5298, 5310, 5318, 5349, 5351, 5353, 5355, 5357, 5358, 5394, 5402, 5405, 5412, 5413, 5417, 5421, 5432, 5433, 5445, 5450, 5457, 5458, 5480, 5481, 5495, 5498, 5499, 5500, 5501, 5517, 5550, 5554, 5555, 5556, 5568, 5601, 5631, 5632, 5656, 5666, 5667, 5670, 5671, 5672, 5678, 5683, 5684, 5693, 5701, 5718, 5719, 5722, 5723, 5724, 5741, 5742, 5800, 5900, 5905, 5931, 5938, 5984, 5985, 5986, 5988, 5989, 6000, 6001, 6002, 6003, 6004, 6005, 6006, 6007, 6008, 6009, 6010, 6011, 6012, 6013, 6014, 6015, 6016, 6017, 6018, 6019, 6020, 6021, 6022, 6023, 6024, 6025, 6026, 6027, 6028, 6029, 6030, 6031, 6032, 6033, 6034, 6035, 6036, 6037, 6038, 6039, 6040, 6041, 6042, 6043, 6044, 6045, 6046, 6047, 6048, 6049, 6050, 6051, 6052, 6053, 6054, 6055, 6056, 6057, 6058, 6059, 6060, 6061, 6062, 6063, 6086, 6100, 6101, 6110, 6111, 6112, 6113, 6136, 6159, 6160, 6161, 6162, 6163, 6164, 6165, 6167, 6170, 6200, 6201, 6225, 6227, 6240, 6244, 6255, 6257, 6260, 6262, 6343, 6346, 6347, 6350, 6379, 6389, 6432, 6436, 6437, 6443, 6444, 6445, 6454, 6463, 6464, 6465, 6466, 6467, 6468, 6469, 6470, 6471, 6472, 6513, 6514, 6515, 6516, 6543, 6556, 6560, 6561, 6566, 6571, 6600, 6601, 6602, 6619, 6622, 6626, 6653, 6660, 6661, 6662, 6663, 6664, 6665, 6666, 6667, 6668, 6669, 6679, 6690, 6697, 6699, 6715, 6771, 6783, 6784, 6785, 6801, 6881, 6882, 6883, 6884, 6885, 6886, 6887, 6888, 6889, 6890, 6891, 6892, 6893, 6894, 6895, 6896, 6897, 6898, 6899, 6900, 6901, 6902, 6903, 6904, 6905, 6906, 6907, 6908, 6909, 6910, 6911, 6912, 6913, 6914, 6915, 6916, 6917, 6918, 6919, 6920, 6921, 6922, 6923, 6924, 6925, 6926, 6927, 6928, 6929, 6930, 6931, 6932, 6933, 6934, 6935, 6936, 6937, 6938, 6939, 6940, 6941, 6942, 6943, 6944, 6945, 6946, 6947, 6948, 6949, 6950, 6951, 6952, 6953, 6954, 6955, 6956, 6957, 6958, 6959, 6960, 6961, 6962, 6963, 6964, 6965, 6966, 6967, 6968, 6969, 6970, 6971, 6972, 6973, 6974, 6975, 6976, 6977, 6978, 6979, 6980, 6981, 6982, 6983, 6984, 6985, 6986, 6987, 6988, 6989, 6990, 6991, 6992, 6993, 6994, 6995, 6996, 6997, 6998, 6999, 7000, 7001, 7002, 7005, 7006, 7010, 7022, 7023, 7025, 7047, 7070, 7077, 7133, 7144, 7145, 7171, 7262, 7272, 7306, 7307, 7312, 7396, 7400, 7401, 7402, 7471, 7473, 7474, 7478, 7542, 7547, 7575, 7624, 7631, 7634, 7652, 7653, 7654, 7655, 7656, 7657, 7658, 7659, 7660, 7670, 7680, 7687, 7707, 7708, 7717, 7777, 7778, 7779, 7780, 7781, 7782, 7783, 7784, 7785, 7786, 7787, 7788, 7831, 7880, 7890, 7915, 7935, 7946, 7979, 7990, 8000, 8005, 8006, 8007, 8008, 8009, 8010, 8042, 8061, 8069, 8070, 8074, 8075, 8080, 8081, 8088, 8089, 8090, 8091, 8092, 8096, 8111, 8112, 8116, 8118, 8123, 8124, 8125, 8139, 8140, 8172, 8184, 8194, 8195, 8200, 8222, 8236, 8243, 8245, 8280, 8281, 8291, 8303, 8332, 8333, 8334, 8337, 8384, 8388, 8400, 8401, 8403, 8443, 8444, 8448, 8484, 8500, 8530, 8531, 8555, 8580, 8611, 8629, 8642, 8691, 8765, 8767, 8834, 8840, 8880, 8883, 8887, 8888, 8889, 8920, 8983, 8997, 8998, 8999, 9000, 9001, 9002, 9003, 9006, 9030, 9042, 9043, 9050, 9051, 9060, 9080, 9081, 9090, 9091, 9092, 9100, 9101, 9102, 9103, 9119, 9150, 9191, 9199, 9200, 9217, 9293, 9295, 9296, 9300, 9303, 9306, 9309, 9312, 9332, 9333, 9339, 9389, 9392, 9418, 9419, 9420, 9421, 9422, 9425, 9443, 9535, 9536, 9600, 9669, 9675, 9676, 9695, 9735, 9785, 9800, 9875, 9897, 9898, 9899, 9901, 9981, 9982, 9987, 9993, 9997, 9999}

func generateRandomPort() int {
	min := 1024
	max := 65000

	for {
		// Generate a random port number between min and max
		port := rand.Intn(max-min+1) + min

		// Check if the port is one of the known ports
		if contains(knownPorts, port) {
			continue
		}

		// Check if the port is in use on localhost
		conn, err := net.Listen("tcp", ":"+strconv.Itoa(port))
		if err == nil {
			conn.Close()
		}
		return port
	}
}

func contains(arr [2035]int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	port := generateRandomPort()
	fmt.Println(port)
}
