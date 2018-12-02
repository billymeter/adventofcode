(ns main)

(use 'clojure.java.io)

(defn get-lines [fname]
  (with-open [r (reader fname)]
  (map #(Integer. %) (doall (line-seq r)))))
;
; (defn get-lines [& params]
;   '(7 7 -2 -7 -4))

(defn part1 []
  (->> (get-lines "input")
       (reduce + 0)
       (println)))

(defn part2 []
  (println
    (reduce (fn [freqs x]
              (let [sum (+ x (first freqs))]
                (if (some #{sum} freqs)
                    (reduced sum)
                    (conj freqs sum))))
            '(0)
            ; (cycle (get-lines "input")))))
            (get-lines "input"))))

(defn -main [& args]
  (do
    (part1)
    (part2)))
