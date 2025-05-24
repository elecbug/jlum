namespace JournalList
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var app = new App();
            app.Run(args);

            Paper paper = new Paper() { Authors = [new Author()] };

            using (var sw = new StreamWriter("paper.json"))
            {
                string json = paper.Serialize();
                sw.Write(json);
            }

            using (var sr = new StreamReader("paper-test.json"))
            {
                string json = sr.ReadToEnd();
                Paper? deserializedPaper = Paper.Deserialize(json);
                if (deserializedPaper != null)
                {
                    Console.WriteLine(deserializedPaper.ToString());
                }
                else
                {
                    Console.WriteLine("Failed to deserialize the paper.");
                }
            }
        }
    }
}