import json
from utils.exec_module import exec_module

module = exec_module("2642. 设计可以求最短路径的图类.py")

def test():
    cases = json.loads("""[
        ["Graph","shortestPath","shortestPath","addEdge","shortestPath"],
            [[4,[[0,2,5],[0,1,2],[1,2,1],[3,0,3]]],[3,2],[0,3],[[1,3,4]],[0,3]],
            [null,6,-1,null,6],
        ["Graph","shortestPath","shortestPath","addEdge","addEdge","shortestPath","addEdge","addEdge","addEdge","addEdge","addEdge","addEdge","addEdge","addEdge","shortestPath","shortestPath","addEdge","addEdge","shortestPath","addEdge","shortestPath","addEdge","shortestPath","addEdge","addEdge","shortestPath","shortestPath","addEdge","shortestPath","addEdge","addEdge","shortestPath","shortestPath","addEdge","shortestPath","addEdge","addEdge","shortestPath","addEdge","addEdge","shortestPath","shortestPath","shortestPath","addEdge","shortestPath","addEdge","shortestPath","addEdge","shortestPath","shortestPath","addEdge","shortestPath","addEdge","addEdge","shortestPath","shortestPath","addEdge","addEdge","shortestPath","addEdge","addEdge","addEdge","addEdge","addEdge","shortestPath","shortestPath"],
            [[17,[[6,3,475731],[3,7,515993],[13,8,904914],[1,15,665336],[3,4,419955],[2,5,591627],[15,13,180374],[8,6,939578],[7,10,459913],[8,1,942800],[14,6,876558],[15,5,215248],[13,14,762427],[1,5,914567],[6,12,919273],[12,16,342212],[10,8,60822],[3,14,947396],[15,0,263172],[10,6,711514],[9,14,120577],[11,5,476839],[11,13,438668],[12,9,527842],[14,16,588402],[15,2,195790],[1,9,785659],[2,7,787223],[11,7,99316],[6,1,948004],[4,12,31559],[5,4,453573],[5,8,141131],[5,12,767697],[1,12,312956],[14,11,374561],[15,11,19433],[0,9,227239],[12,10,325409],[16,13,413278],[10,1,155788],[5,3,294209],[7,5,54490],[3,13,701716],[2,8,927178],[12,14,367956],[14,10,953761],[3,16,9061],[2,3,421223],[4,15,189155],[14,2,711954],[16,15,734202],[7,4,917325],[0,1,818496],[8,3,548282],[16,12,385482]]],[0,11],[14,14],[[9,2,3168]],[[16,5,261919]],[6,7],[[1,13,2527]],[[7,0,1564]],[[16,1,1359]],[[5,14,213]],[[15,10,31158]],[[6,2,144]],[[6,9,84]],[[5,15,957477]],[12,2],[13,1],[[15,4,71]],[[2,14,48]],[8,0],[[9,11,568935]],[8,5],[[3,8,674529]],[8,10],[[2,16,38]],[[11,16,11]],[3,5],[14,5],[[15,12,9]],[16,8],[[9,4,69863]],[[0,12,1]],[14,6],[3,2],[[0,11,987029]],[12,2],[[6,7,1]],[[7,1,43668]],[9,14],[[3,10,1]],[[2,10,1]],[15,12],[11,10],[16,4],[[1,4,613267]],[6,16],[[10,5,1]],[9,1],[[15,14,1]],[5,1],[13,16],[[4,6,654722]],[3,14],[[6,10,1]],[[14,7,1]],[3,10],[9,15],[[14,12,454433]],[[5,2,1]],[12,5],[[16,10,1]],[[15,6,1]],[[11,3,1]],[[10,16,930106]],[[6,14,1]],[6,14],[8,9]],
            [null,722377,0,null,null,991724,null,null,null,null,null,null,null,null,531010,1352188,null,null,1065839,null,819262,null,1188550,null,null,270980,528367,null,403050,null,null,876558,747964,null,531010,null,null,3216,null,null,9,331288,666766,null,182,null,4565,null,304629,1136999,null,215,null,null,1,259018,null,null,325410,null,null,null,null,null,1,777137]
    ]""")
    us = 3
    cnt = len(cases) // us
    for s in [
        module.Graph,
        module.Graph_2,
        module.Graph_3,
    ]:
        sname = s.__class__.__name__
        # try:
        g = None
        for i in range(cnt):
            data = cases[us*i : us*(i+1)-1]
            want = cases[us*(i+1)-1]
            n = len(data[0])
            ans = [None] * n
            for j, m in enumerate(data[0]):
                match m:
                    case "Graph":
                        g = s(*data[1][j])
                        ans[j] = None
                    case "addEdge":
                        ans[j] = g.addEdge(*data[1][j])
                    case "shortestPath":
                        ans[j] = g.shortestPath(*data[1][j])
            assert ans == want, sname
        # except:
        #     assert False, sname
