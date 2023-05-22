from typing import *
class Solution:
    def displayTable(self, orders: List[List[str]]) -> List[List[str]]:
        dic = DefaultDict(List)
        for order in orders:
            if order[1] in dic:
                dic[order[1]].append(order[2])
            else:
                dic[order[1]] = [order[2]]
        
        menu = [['Table']]
        dish = set()
        for v in dic.values():
            for i in range(len(v)):
                dish.add(v[i])
        dish = list(dish)
        dish.sort()
        for dis in dish:
            menu[0].append(dis)
        for k in sorted(dic,key = lambda x: int(x)):
            temp = [k]
            for d in dish:
                temp.append(str(dic[k].count(d)))
            menu.append(temp)
        return menu

def main():
    lst = [["David","3","Ceviche"],["Corina","10","Beef Burrito"],["David","3","Fried Chicken"],["Carla","5","Water"],["Carla","5","Ceviche"],["Rous","3","Ceviche"]]
    s = Solution()
    print(s.displayTable(lst))
main()