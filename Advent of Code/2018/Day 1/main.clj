(ns main)

(use 'clojure.java.io)

(defn get-lines [fname]
  (with-open [r (reader fname)]
  (map #(Integer. %) (doall (line-seq r)))))

(defn part1 []
  (->> (get-lines "input")
       (reduce + 0)
       (println)))

(defn part2 []
  "Brute forces the answer. It take a few minutes to complete."
  (println
    (reduce (fn [freqs x]
              (let [sum (+ x (first freqs))]
                (if (some #{sum} freqs)
                    (reduced sum)
                    (conj freqs sum))))
            '(0)
            (cycle (get-lines "input")))))

(defn -main [& args]
  (do
    (part1)))
    ; (part1)
    ; (part2)))
