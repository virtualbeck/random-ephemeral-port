package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"sync"
)

var knownPorts = map[int]bool{
	10000: true, 10001: true, 10002: true, 10003: true, 10004: true, 10005: true, 10006: true, 10007: true, 10008: true, 10009: true, 10010: true, 10011: true, 10012: true, 10013: true, 10014: true, 10015: true, 10016: true, 10017: true, 10018: true, 10019: true, 10020: true, 10021: true, 10022: true, 10023: true, 10024: true, 10025: true, 10026: true, 10027: true, 10028: true, 10029: true, 10030: true, 10031: true, 10032: true, 10033: true, 10034: true, 10035: true, 10036: true, 10037: true, 10038: true, 10039: true, 10040: true, 10041: true, 10042: true, 10043: true, 10044: true, 10045: true, 10046: true, 10047: true, 10048: true, 10049: true, 10050: true, 10051: true, 10052: true, 10053: true, 10054: true, 10055: true, 10056: true, 10057: true, 10058: true, 10059: true, 10060: true, 10061: true, 10062: true, 10063: true, 10064: true, 10065: true, 10066: true, 10067: true, 10068: true, 10069: true, 10070: true, 10071: true, 10072: true, 10073: true, 10074: true, 10075: true, 10076: true, 10077: true, 10078: true, 10079: true, 10080: true, 10081: true, 10082: true, 10083: true, 10084: true, 10085: true, 10086: true, 10087: true, 10088: true, 10089: true, 10090: true, 10091: true, 10092: true, 10093: true, 10094: true, 10095: true, 10096: true, 10097: true, 10098: true, 10099: true, 10100: true, 10101: true, 10102: true, 10103: true, 10104: true, 10105: true, 10106: true, 10107: true, 10108: true, 10109: true, 10110: true, 10111: true, 10112: true, 10113: true, 10114: true, 10115: true, 10116: true, 10117: true, 10118: true, 10119: true, 10120: true, 10121: true, 10122: true, 10123: true, 10124: true, 10125: true, 10126: true, 10127: true, 10128: true, 10129: true, 10130: true, 10131: true, 10132: true, 10133: true, 10134: true, 10135: true, 10136: true, 10137: true, 10138: true, 10139: true, 10140: true, 10141: true, 10142: true, 10143: true, 10144: true, 10145: true, 10146: true, 10147: true, 10148: true, 10149: true, 10150: true, 10151: true, 10152: true, 10153: true, 10154: true, 10155: true, 10156: true, 10157: true, 10158: true, 10159: true, 10160: true, 10161: true, 10162: true, 10163: true, 10164: true, 10165: true, 10166: true, 10167: true, 10168: true, 10169: true, 10170: true, 10171: true, 10172: true, 10173: true, 10174: true, 10175: true, 10176: true, 10177: true, 10178: true, 10179: true, 10180: true, 10181: true, 10182: true, 10183: true, 10184: true, 10185: true, 10186: true, 10187: true, 10188: true, 10189: true, 10190: true, 10191: true, 10192: true, 10193: true, 10194: true, 10195: true, 10196: true, 10197: true, 10198: true, 10199: true, 10200: true, 10201: true, 10202: true, 10203: true, 10204: true, 10212: true, 1024: true, 1025: true, 1027: true, 1028: true, 1029: true, 10308: true, 10346: true, 10468: true, 10480: true, 10505: true, 10514: true, 10578: true, 1058: true, 1059: true, 1080: true, 10800: true, 10823: true, 10836: true, 1085: true, 10891: true, 10933: true, 1098: true, 1099: true, 11001: true, 1109: true, 11100: true, 11111: true, 11112: true, 1112: true, 1113: true, 1119: true, 11211: true, 11214: true, 11215: true, 11235: true, 11311: true, 11371: true, 1167: true, 11753: true, 1194: true, 1198: true, 12000: true, 12012: true, 12013: true, 12035: true, 12043: true, 12046: true, 1212: true, 1214: true, 1220: true, 12201: true, 12222: true, 12223: true, 12307: true, 12308: true, 1234: true, 12345: true, 1241: true, 12443: true, 12489: true, 1270: true, 1293: true, 12975: true, 13000: true, 13001: true, 13002: true, 13003: true, 13004: true, 13005: true, 13006: true, 13007: true, 13008: true, 13009: true, 13010: true, 13011: true, 13012: true, 13013: true, 13014: true, 13015: true, 13016: true, 13017: true, 13018: true, 13019: true, 13020: true, 13021: true, 13022: true, 13023: true, 13024: true, 13025: true, 13026: true, 13027: true, 13028: true, 13029: true, 13030: true, 13031: true, 13032: true, 13033: true, 13034: true, 13035: true, 13036: true, 13037: true, 13038: true, 13039: true, 13040: true, 13041: true, 13042: true, 13043: true, 13044: true, 13045: true, 13046: true, 13047: true, 13048: true, 13049: true, 13050: true, 13075: true, 1311: true, 1314: true, 1319: true, 1337: true, 13400: true, 1341: true, 1344: true, 1352: true, 1360: true, 13698: true, 13720: true, 13721: true, 13724: true, 13782: true, 13783: true, 13785: true, 13786: true, 1414: true, 1417: true, 1418: true, 1419: true, 1420: true, 1431: true, 1433: true, 1434: true, 14550: true, 14567: true, 14652: true, 1476: true, 14800: true, 1481: true, 1492: true, 1494: true, 1500: true, 15000: true, 15009: true, 1501: true, 15010: true, 1503: true, 1512: true, 1513: true, 1521: true, 1524: true, 1527: true, 1533: true, 1534: true, 15345: true, 1540: true, 1541: true, 1542: true, 15441: true, 1545: true, 1547: true, 1550: true, 15567: true, 1560: true, 1561: true, 1562: true, 1563: true, 1564: true, 1565: true, 1566: true, 1567: true, 15672: true, 1568: true, 1569: true, 1570: true, 1571: true, 1572: true, 1573: true, 1574: true, 1575: true, 1576: true, 1577: true, 1578: true, 1579: true, 1580: true, 1581: true, 1582: true, 1583: true, 1584: true, 1585: true, 1586: true, 1587: true, 1588: true, 1589: true, 1590: true, 16000: true, 1604: true, 16080: true, 16200: true, 16225: true, 16250: true, 1626: true, 16261: true, 1627: true, 1628: true, 1629: true, 16300: true, 16384: true, 16385: true, 16386: true, 16387: true, 16393: true, 16394: true, 16395: true, 16396: true, 16397: true, 16398: true, 16399: true, 16400: true, 16401: true, 16402: true, 16403: true, 16404: true, 16405: true, 16406: true, 16407: true, 16408: true, 16409: true, 16410: true, 16411: true, 16412: true, 16413: true, 16414: true, 16415: true, 16416: true, 16417: true, 16418: true, 16419: true, 16420: true, 16421: true, 16422: true, 16423: true, 16424: true, 16425: true, 16426: true, 16427: true, 16428: true, 16429: true, 16430: true, 16431: true, 16432: true, 16433: true, 16434: true, 16435: true, 16436: true, 16437: true, 16438: true, 16439: true, 16440: true, 16441: true, 16442: true, 16443: true, 16444: true, 16445: true, 16446: true, 16447: true, 16448: true, 16449: true, 1645: true, 16450: true, 16451: true, 16452: true, 16453: true, 16454: true, 16455: true, 16456: true, 16457: true, 16458: true, 16459: true, 1646: true, 16460: true, 16461: true, 16462: true, 16463: true, 16464: true, 16465: true, 16466: true, 16467: true, 16468: true, 16469: true, 16470: true, 16471: true, 16472: true, 16567: true, 1666: true, 16666: true, 16677: true, 1677: true, 1688: true, 17000: true, 1701: true, 17011: true, 1707: true, 1714: true, 1715: true, 1716: true, 1717: true, 1718: true, 1719: true, 1720: true, 1721: true, 1722: true, 17224: true, 17225: true, 1723: true, 1724: true, 1725: true, 1726: true, 1727: true, 1728: true, 1729: true, 1730: true, 1731: true, 1732: true, 1733: true, 17333: true, 1734: true, 1735: true, 1736: true, 1737: true, 1738: true, 1739: true, 1740: true, 1741: true, 1742: true, 1743: true, 1744: true, 1745: true, 1746: true, 1747: true, 17472: true, 17474: true, 17475: true, 1748: true, 1749: true, 1750: true, 17500: true, 1751: true, 1752: true, 1753: true, 1754: true, 1755: true, 1756: true, 1757: true, 1758: true, 1759: true, 1760: true, 1761: true, 1762: true, 1763: true, 1764: true, 1776: true, 17777: true, 1783: true, 1801: true, 18080: true, 18081: true, 18091: true, 18092: true, 18104: true, 1812: true, 1813: true, 18200: true, 18201: true, 18206: true, 18300: true, 18301: true, 18306: true, 18333: true, 18400: true, 18401: true, 18505: true, 18506: true, 18605: true, 18606: true, 1863: true, 18676: true, 1880: true, 1883: true, 1900: true, 19000: true, 19001: true, 19132: true, 19133: true, 19150: true, 19226: true, 19294: true, 19295: true, 19302: true, 1935: true, 19531: true, 19532: true, 1965: true, 1967: true, 1972: true, 19788: true, 19812: true, 19813: true, 19814: true, 1984: true, 1985: true, 1998: true, 19999: true, 2000: true, 20000: true, 2010: true, 2033: true, 2049: true, 2056: true, 20560: true, 20582: true, 20583: true, 20595: true, 2080: true, 20808: true, 2082: true, 2083: true, 2086: true, 2087: true, 2095: true, 2096: true, 2100: true, 2101: true, 2102: true, 21025: true, 2103: true, 2104: true, 21064: true, 2123: true, 2142: true, 2152: true, 2159: true, 2181: true, 2195: true, 2196: true, 2197: true, 22000: true, 2210: true, 2211: true, 22136: true, 2221: true, 2222: true, 22222: true, 2223: true, 2224: true, 2225: true, 2226: true, 22347: true, 22350: true, 22351: true, 2240: true, 2261: true, 2262: true, 2302: true, 2303: true, 2305: true, 23073: true, 23399: true, 2351: true, 23513: true, 2368: true, 2369: true, 2370: true, 2372: true, 2375: true, 2376: true, 2377: true, 2379: true, 2380: true, 2389: true, 2399: true, 2401: true, 2404: true, 2424: true, 2427: true, 24441: true, 24444: true, 24465: true, 2447: true, 24554: true, 2456: true, 2459: true, 2480: true, 24800: true, 2483: true, 2484: true, 24842: true, 2500: true, 2501: true, 2535: true, 2541: true, 2546: true, 2547: true, 2548: true, 25565: true, 25575: true, 25600: true, 25734: true, 25826: true, 2593: true, 2598: true, 2599: true, 26000: true, 2628: true, 2638: true, 26822: true, 26900: true, 26901: true, 26909: true, 26910: true, 26911: true, 27000: true, 27001: true, 27002: true, 27003: true, 27004: true, 27005: true, 27006: true, 27007: true, 27008: true, 27009: true, 27015: true, 27016: true, 27017: true, 27018: true, 27019: true, 27020: true, 27021: true, 27022: true, 27023: true, 27024: true, 27025: true, 27026: true, 27027: true, 27028: true, 27029: true, 27030: true, 27031: true, 27032: true, 27033: true, 27034: true, 27035: true, 27036: true, 2710: true, 2727: true, 27374: true, 27500: true, 27501: true, 27502: true, 27503: true, 27504: true, 27505: true, 27506: true, 27507: true, 27508: true, 27509: true, 27510: true, 27511: true, 27512: true, 27513: true, 27514: true, 27515: true, 27516: true, 27517: true, 27518: true, 27519: true, 27520: true, 27521: true, 27522: true, 27523: true, 27524: true, 27525: true, 27526: true, 27527: true, 27528: true, 27529: true, 27530: true, 27531: true, 27532: true, 27533: true, 27534: true, 27535: true, 27536: true, 27537: true, 27538: true, 27539: true, 27540: true, 27541: true, 27542: true, 27543: true, 27544: true, 27545: true, 27546: true, 27547: true, 27548: true, 27549: true, 27550: true, 27551: true, 27552: true, 27553: true, 27554: true, 27555: true, 27556: true, 27557: true, 27558: true, 27559: true, 27560: true, 27561: true, 27562: true, 27563: true, 27564: true, 27565: true, 27566: true, 27567: true, 27568: true, 27569: true, 27570: true, 27571: true, 27572: true, 27573: true, 27574: true, 27575: true, 27576: true, 27577: true, 27578: true, 27579: true, 27580: true, 27581: true, 27582: true, 27583: true, 27584: true, 27585: true, 27586: true, 27587: true, 27588: true, 27589: true, 2759: true, 27590: true, 27591: true, 27592: true, 27593: true, 27594: true, 27595: true, 27596: true, 27597: true, 27598: true, 27599: true, 27600: true, 27601: true, 27602: true, 27603: true, 27604: true, 27605: true, 27606: true, 27607: true, 27608: true, 27609: true, 27610: true, 27611: true, 27612: true, 27613: true, 27614: true, 27615: true, 27616: true, 27617: true, 27618: true, 27619: true, 27620: true, 27621: true, 27622: true, 27623: true, 27624: true, 27625: true, 27626: true, 27627: true, 27628: true, 27629: true, 27630: true, 27631: true, 27632: true, 27633: true, 27634: true, 27635: true, 27636: true, 27637: true, 27638: true, 27639: true, 27640: true, 27641: true, 27642: true, 27643: true, 27644: true, 27645: true, 27646: true, 27647: true, 27648: true, 27649: true, 27650: true, 27651: true, 27652: true, 27653: true, 27654: true, 27655: true, 27656: true, 27657: true, 27658: true, 27659: true, 27660: true, 27661: true, 27662: true, 27663: true, 27664: true, 27665: true, 27666: true, 27667: true, 27668: true, 27669: true, 27670: true, 27671: true, 27672: true, 27673: true, 27674: true, 27675: true, 27676: true, 27677: true, 27678: true, 27679: true, 27680: true, 27681: true, 27682: true, 27683: true, 27684: true, 27685: true, 27686: true, 27687: true, 27688: true, 27689: true, 27690: true, 27691: true, 27692: true, 27693: true, 27694: true, 27695: true, 27696: true, 27697: true, 27698: true, 27699: true, 27700: true, 27701: true, 27702: true, 27703: true, 27704: true, 27705: true, 27706: true, 27707: true, 27708: true, 27709: true, 27710: true, 27711: true, 27712: true, 27713: true, 27714: true, 27715: true, 27716: true, 27717: true, 27718: true, 27719: true, 27720: true, 27721: true, 27722: true, 27723: true, 27724: true, 27725: true, 27726: true, 27727: true, 27728: true, 27729: true, 27730: true, 27731: true, 27732: true, 27733: true, 27734: true, 27735: true, 27736: true, 27737: true, 27738: true, 27739: true, 27740: true, 27741: true, 27742: true, 27743: true, 27744: true, 27745: true, 27746: true, 27747: true, 27748: true, 27749: true, 2775: true, 27750: true, 27751: true, 27752: true, 27753: true, 27754: true, 27755: true, 27756: true, 27757: true, 27758: true, 27759: true, 27760: true, 27761: true, 27762: true, 27763: true, 27764: true, 27765: true, 27766: true, 27767: true, 27768: true, 27769: true, 27770: true, 27771: true, 27772: true, 27773: true, 27774: true, 27775: true, 27776: true, 27777: true, 27778: true, 27779: true, 27780: true, 27781: true, 27782: true, 27783: true, 27784: true, 27785: true, 27786: true, 27787: true, 27788: true, 27789: true, 27790: true, 27791: true, 27792: true, 27793: true, 27794: true, 27795: true, 27796: true, 27797: true, 27798: true, 27799: true, 27800: true, 27801: true, 27802: true, 27803: true, 27804: true, 27805: true, 27806: true, 27807: true, 27808: true, 27809: true, 27810: true, 27811: true, 27812: true, 27813: true, 27814: true, 27815: true, 27816: true, 27817: true, 27818: true, 27819: true, 27820: true, 27821: true, 27822: true, 27823: true, 27888: true, 27901: true, 27902: true, 27903: true, 27904: true, 27905: true, 27906: true, 27907: true, 27908: true, 27909: true, 27910: true, 27950: true, 27960: true, 27961: true, 27962: true, 27963: true, 27964: true, 27965: true, 27966: true, 27967: true, 27968: true, 27969: true, 28000: true, 28001: true, 28015: true, 28016: true, 2809: true, 2811: true, 28200: true, 28260: true, 2827: true, 28443: true, 28769: true, 28770: true, 28771: true, 28785: true, 28786: true, 28852: true, 28910: true, 28960: true, 29000: true, 29070: true, 2944: true, 2945: true, 2947: true, 2948: true, 2949: true, 2967: true, 2989: true, 29900: true, 29901: true, 29920: true, 3000: true, 30000: true, 30003: true, 30004: true, 3001: true, 30033: true, 3004: true, 3010: true, 30120: true, 3020: true, 3050: true, 3052: true, 30564: true, 3074: true, 3101: true, 3128: true, 31337: true, 31416: true, 31438: true, 31457: true, 32137: true, 3225: true, 3233: true, 32400: true, 3260: true, 3268: true, 3269: true, 32764: true, 3283: true, 32887: true, 3290: true, 32976: true, 3305: true, 3306: true, 3323: true, 3332: true, 3333: true, 33434: true, 3344: true, 3351: true, 33848: true, 3386: true, 3389: true, 3396: true, 34000: true, 3412: true, 34197: true, 3423: true, 3424: true, 3435: true, 3455: true, 3478: true, 3479: true, 3480: true, 3483: true, 3493: true, 3503: true, 3516: true, 3527: true, 3535: true, 35357: true, 3544: true, 3551: true, 3601: true, 3632: true, 36330: true, 3645: true, 3655: true, 3659: true, 3667: true, 3671: true, 3689: true, 3690: true, 37008: true, 3702: true, 3721: true, 3724: true, 3725: true, 3749: true, 3768: true, 3784: true, 3785: true, 3799: true, 3804: true, 3825: true, 3826: true, 3830: true, 3835: true, 3856: true, 3868: true, 3872: true, 3880: true, 3900: true, 3960: true, 3962: true, 3978: true, 3979: true, 3999: true, 4000: true, 40000: true, 4001: true, 4018: true, 4035: true, 4045: true, 4050: true, 4069: true, 4070: true, 4089: true, 4090: true, 4093: true, 4096: true, 4105: true, 4111: true, 41121: true, 4116: true, 4123: true, 41230: true, 4125: true, 4172: true, 41794: true, 41795: true, 41796: true, 41797: true, 4190: true, 4195: true, 4197: true, 4198: true, 4200: true, 4201: true, 42081: true, 4222: true, 4226: true, 4242: true, 4243: true, 4244: true, 42590: true, 42999: true, 4303: true, 4307: true, 43110: true, 4321: true, 43594: true, 43595: true, 44123: true, 44405: true, 4444: true, 4445: true, 44818: true, 4486: true, 4488: true, 4500: true, 4502: true, 4503: true, 4504: true, 4505: true, 4506: true, 4534: true, 4560: true, 4567: true, 4569: true, 4604: true, 4605: true, 4610: true, 4611: true, 4612: true, 4613: true, 4614: true, 4615: true, 4616: true, 4617: true, 4618: true, 4619: true, 4620: true, 4621: true, 4622: true, 4623: true, 4624: true, 4625: true, 4626: true, 4627: true, 4628: true, 4629: true, 4630: true, 4631: true, 4632: true, 4633: true, 4634: true, 4635: true, 4636: true, 4637: true, 4638: true, 4639: true, 4640: true, 4662: true, 4664: true, 4672: true, 4711: true, 4713: true, 4723: true, 4724: true, 4728: true, 4730: true, 4739: true, 4747: true, 4753: true, 4757: true, 47808: true, 47809: true, 47810: true, 47811: true, 47812: true, 47813: true, 47814: true, 47815: true, 47816: true, 47817: true, 47818: true, 47819: true, 47820: true, 47821: true, 47822: true, 47823: true, 4789: true, 4791: true, 4840: true, 4843: true, 4847: true, 4848: true, 48556: true, 48656: true, 48657: true, 4894: true, 49151: true, 4944: true, 4949: true, 4950: true, 5000: true, 5001: true, 5002: true, 5003: true, 5004: true, 5005: true, 5006: true, 5007: true, 5008: true, 5009: true, 5010: true, 5011: true, 5012: true, 5013: true, 5014: true, 5015: true, 5016: true, 5017: true, 5018: true, 5019: true, 5020: true, 5021: true, 5022: true, 5023: true, 5024: true, 5025: true, 5026: true, 5027: true, 5028: true, 5029: true, 5030: true, 5031: true, 5032: true, 5033: true, 5034: true, 5035: true, 5036: true, 5037: true, 5038: true, 5039: true, 5040: true, 5041: true, 5042: true, 5043: true, 5044: true, 5045: true, 5046: true, 5047: true, 5048: true, 5049: true, 5050: true, 5051: true, 5052: true, 5053: true, 5054: true, 5055: true, 5056: true, 5057: true, 5058: true, 5059: true, 5060: true, 5061: true, 5062: true, 5063: true, 5064: true, 5065: true, 5070: true, 5080: true, 5084: true, 5085: true, 5090: true, 5093: true, 5099: true, 5104: true, 5121: true, 5124: true, 5125: true, 5150: true, 5151: true, 5154: true, 5172: true, 5173: true, 5190: true, 5198: true, 5199: true, 5200: true, 5201: true, 5222: true, 5223: true, 5228: true, 5231: true, 5232: true, 5235: true, 5242: true, 5243: true, 5246: true, 5247: true, 5269: true, 5280: true, 5281: true, 5298: true, 5310: true, 5318: true, 5349: true, 5351: true, 5353: true, 5355: true, 5357: true, 5358: true, 5394: true, 5402: true, 5405: true, 5412: true, 5413: true, 5417: true, 5421: true, 5432: true, 5433: true, 5445: true, 5450: true, 5457: true, 5458: true, 5480: true, 5481: true, 5495: true, 5498: true, 5499: true, 5500: true, 5501: true, 5517: true, 5550: true, 5554: true, 5555: true, 5556: true, 5568: true, 5601: true, 5631: true, 5632: true, 5656: true, 5666: true, 5667: true, 5670: true, 5671: true, 5672: true, 5678: true, 5683: true, 5684: true, 5693: true, 5701: true, 5718: true, 5719: true, 5722: true, 5723: true, 5724: true, 5741: true, 5742: true, 5800: true, 5900: true, 5905: true, 5931: true, 5938: true, 5984: true, 5985: true, 5986: true, 5988: true, 5989: true, 6000: true, 6001: true, 6002: true, 6003: true, 6004: true, 6005: true, 6006: true, 6007: true, 6008: true, 6009: true, 6010: true, 6011: true, 6012: true, 6013: true, 6014: true, 6015: true, 6016: true, 6017: true, 6018: true, 6019: true, 6020: true, 6021: true, 6022: true, 6023: true, 6024: true, 6025: true, 6026: true, 6027: true, 6028: true, 6029: true, 6030: true, 6031: true, 6032: true, 6033: true, 6034: true, 6035: true, 6036: true, 6037: true, 6038: true, 6039: true, 6040: true, 6041: true, 6042: true, 6043: true, 6044: true, 6045: true, 6046: true, 6047: true, 6048: true, 6049: true, 6050: true, 6051: true, 6052: true, 6053: true, 6054: true, 6055: true, 6056: true, 6057: true, 6058: true, 6059: true, 6060: true, 6061: true, 6062: true, 6063: true, 6086: true, 6100: true, 6101: true, 6110: true, 6111: true, 6112: true, 6113: true, 6136: true, 6159: true, 6160: true, 6161: true, 6162: true, 6163: true, 6164: true, 6165: true, 6167: true, 6170: true, 6200: true, 6201: true, 6225: true, 6227: true, 6240: true, 6244: true, 6255: true, 6257: true, 6260: true, 6262: true, 6343: true, 6346: true, 6347: true, 6350: true, 6379: true, 6389: true, 6432: true, 6436: true, 6437: true, 6443: true, 6444: true, 6445: true, 6454: true, 6463: true, 6464: true, 6465: true, 6466: true, 6467: true, 6468: true, 6469: true, 6470: true, 6471: true, 6472: true, 6513: true, 6514: true, 6515: true, 6516: true, 6543: true, 6556: true, 6560: true, 6561: true, 6566: true, 6571: true, 6600: true, 6601: true, 6602: true, 6619: true, 6622: true, 6626: true, 6653: true, 6660: true, 6661: true, 6662: true, 6663: true, 6664: true, 6665: true, 6666: true, 6667: true, 6668: true, 6669: true, 6679: true, 6690: true, 6697: true, 6699: true, 6715: true, 6771: true, 6783: true, 6784: true, 6785: true, 6801: true, 6881: true, 6882: true, 6883: true, 6884: true, 6885: true, 6886: true, 6887: true, 6888: true, 6889: true, 6890: true, 6891: true, 6892: true, 6893: true, 6894: true, 6895: true, 6896: true, 6897: true, 6898: true, 6899: true, 6900: true, 6901: true, 6902: true, 6903: true, 6904: true, 6905: true, 6906: true, 6907: true, 6908: true, 6909: true, 6910: true, 6911: true, 6912: true, 6913: true, 6914: true, 6915: true, 6916: true, 6917: true, 6918: true, 6919: true, 6920: true, 6921: true, 6922: true, 6923: true, 6924: true, 6925: true, 6926: true, 6927: true, 6928: true, 6929: true, 6930: true, 6931: true, 6932: true, 6933: true, 6934: true, 6935: true, 6936: true, 6937: true, 6938: true, 6939: true, 6940: true, 6941: true, 6942: true, 6943: true, 6944: true, 6945: true, 6946: true, 6947: true, 6948: true, 6949: true, 6950: true, 6951: true, 6952: true, 6953: true, 6954: true, 6955: true, 6956: true, 6957: true, 6958: true, 6959: true, 6960: true, 6961: true, 6962: true, 6963: true, 6964: true, 6965: true, 6966: true, 6967: true, 6968: true, 6969: true, 6970: true, 6971: true, 6972: true, 6973: true, 6974: true, 6975: true, 6976: true, 6977: true, 6978: true, 6979: true, 6980: true, 6981: true, 6982: true, 6983: true, 6984: true, 6985: true, 6986: true, 6987: true, 6988: true, 6989: true, 6990: true, 6991: true, 6992: true, 6993: true, 6994: true, 6995: true, 6996: true, 6997: true, 6998: true, 6999: true, 7000: true, 7001: true, 7002: true, 7005: true, 7006: true, 7010: true, 7022: true, 7023: true, 7025: true, 7047: true, 7070: true, 7077: true, 7133: true, 7144: true, 7145: true, 7171: true, 7262: true, 7272: true, 7306: true, 7307: true, 7312: true, 7396: true, 7400: true, 7401: true, 7402: true, 7471: true, 7473: true, 7474: true, 7478: true, 7542: true, 7547: true, 7575: true, 7624: true, 7631: true, 7634: true, 7652: true, 7653: true, 7654: true, 7655: true, 7656: true, 7657: true, 7658: true, 7659: true, 7660: true, 7670: true, 7680: true, 7687: true, 7707: true, 7708: true, 7717: true, 7777: true, 7778: true, 7779: true, 7780: true, 7781: true, 7782: true, 7783: true, 7784: true, 7785: true, 7786: true, 7787: true, 7788: true, 7831: true, 7880: true, 7890: true, 7915: true, 7935: true, 7946: true, 7979: true, 7990: true, 8000: true, 8005: true, 8006: true, 8007: true, 8008: true, 8009: true, 8010: true, 8042: true, 8061: true, 8069: true, 8070: true, 8074: true, 8075: true, 8080: true, 8081: true, 8088: true, 8089: true, 8090: true, 8091: true, 8092: true, 8096: true, 8111: true, 8112: true, 8116: true, 8118: true, 8123: true, 8124: true, 8125: true, 8139: true, 8140: true, 8172: true, 8184: true, 8194: true, 8195: true, 8200: true, 8222: true, 8236: true, 8243: true, 8245: true, 8280: true, 8281: true, 8291: true, 8303: true, 8332: true, 8333: true, 8334: true, 8337: true, 8384: true, 8388: true, 8400: true, 8401: true, 8403: true, 8443: true, 8444: true, 8448: true, 8484: true, 8500: true, 8530: true, 8531: true, 8555: true, 8580: true, 8611: true, 8629: true, 8642: true, 8691: true, 8765: true, 8767: true, 8834: true, 8840: true, 8880: true, 8883: true, 8887: true, 8888: true, 8889: true, 8920: true, 8983: true, 8997: true, 8998: true, 8999: true, 9000: true, 9001: true, 9002: true, 9003: true, 9006: true, 9030: true, 9042: true, 9043: true, 9050: true, 9051: true, 9060: true, 9080: true, 9081: true, 9090: true, 9091: true, 9092: true, 9100: true, 9101: true, 9102: true, 9103: true, 9119: true, 9150: true, 9191: true, 9199: true, 9200: true, 9217: true, 9293: true, 9295: true, 9296: true, 9300: true, 9303: true, 9306: true, 9309: true, 9312: true, 9332: true, 9333: true, 9339: true, 9389: true, 9392: true, 9418: true, 9419: true, 9420: true, 9421: true, 9422: true, 9425: true, 9443: true, 9535: true, 9536: true, 9600: true, 9669: true, 9675: true, 9676: true, 9695: true, 9735: true, 9785: true, 9800: true, 9875: true, 9897: true, 9898: true, 9899: true, 9901: true, 9981: true, 9982: true, 9987: true, 9993: true, 9997: true, 9999: true,
}

func generateRandomPorts(min, max, numPorts int) []int {
	ports := make([]int, numPorts)

	for i := 0; i < numPorts; i++ {
		ports[i] = rand.Intn(max-min+1) + min
	}

	return ports
}

func checkPortsAvailability(ports []int, wg *sync.WaitGroup, availablePorts chan<- int) {
	defer wg.Done()

	for _, port := range ports {
		// Check if the port is in use on localhost
		conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			conn.Close()

			// Check if the port is known
			_, portKnown := knownPorts[port]
			if !portKnown {
				availablePorts <- port
			}
		}
	}
}

func main() {

	min := 1024
	max := 65000
	numPorts := len(knownPorts)
	numWorkers := 8

	// Generate random port candidates
	ports := generateRandomPorts(min, max, numPorts)

	// Create a channel to receive the available port
	availablePorts := make(chan int)

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start goroutines to check ports availability
	for i := 0; i < numWorkers; i++ {
		start := i * (numPorts / numWorkers)
		end := (i + 1) * (numPorts / numWorkers)

		// Adjust the last worker for the remaining ports
		if i == numWorkers-1 {
			end = numPorts
		}

		go checkPortsAvailability(ports[start:end], &wg, availablePorts)
	}

	go func() {
		wg.Wait()
		close(availablePorts)
	}()

	// Get the first available port from the channel
	port := <-availablePorts

	fmt.Println(port)
	os.Exit(0)
}
