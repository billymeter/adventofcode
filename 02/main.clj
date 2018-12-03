(ns main)

(use 'clojure.java.io)

(defn get-lines []
  (with-open [r (reader "input")]
    (doall (line-seq r))))

(defn two? [coll]
  (count (select-keys coll (for [[k v] coll :when (= v 2)] k))))

(defn three? [coll]
  (count (select-keys coll (for [[k v] coll :when (= v 3)] k))))

(defn part1 []
  (let [freqs (map frequencies (get-lines))]
    (*
      (count (filter pos? (map two? freqs)))
      (count (filter pos? (map three? freqs))))))

(defn hamming [x y]
  (count (filter true? (map (partial reduce not=) (map vector x y)))))


(defn part2 []
  (let [lines (get-lines)]
    (map (partial hamming (first lines)) lines)))

(defn -main [& args]
  (println
    (part1)
    (part2)))
