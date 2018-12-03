(ns main)

(use 'clojure.java.io)

(defn get-lines [fname]
  (with-open [r (reader fname)]
    (doall (line-seq r))))

(defn two? [coll]
  (count (select-keys coll (for [[k v] coll :when (= v 2)] k))))

(defn three? [coll]
  (count (select-keys coll (for [[k v] coll :when (= v 3)] k))))

(defn part1 []
  (let [freqs (map frequencies (get-lines "input"))]
    (*
      (count (filter pos? (map two? freqs)))
      (count (filter pos? (map three? freqs))))))

(defn part2 []
  1)

(defn -main [& args]
  (println
    (part1)
    (part2)))
