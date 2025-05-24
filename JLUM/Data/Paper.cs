using System.Text.Json;

namespace JLUM.Data
{
    public struct Paper
    {
        public List<Author> Authors { get; set; }
        public string Title { get; set; }
        public Information Information { get; set; }

        public string Abstract { get; set; }
        public string[] Keywords { get; set; }
        public string[] Tags { get; set; }
        public string[] References { get; set; }
        public string Note { get; set; }
        public string FilePath { get; set; }

        public Languages Language { get; set; }
        public string JournalURL { get; set; }

        public static Paper Empty => new Paper
        {
            Authors = [],
            Title = "",
            Information = Information.Empty,
            Abstract = "",
            Keywords = [],
            Tags = [],
            References = [],
            Note = "",
            FilePath = "",
            Language = Languages.Unknown,
            JournalURL = ""
        };

        public override string ToString()
        {
            if (Authors.Count == 0)
                Authors.Add(new Author() { Name = "[WARNING]DO_NOT_FIND_AUTHORS" });

            List<Author> authors = Authors[0..(Authors.Count - 1)];
            Author lastAuthor = Authors[Authors.Count - 1];
            string authorString = string.Join(", ", authors.Select(a => a.ToString()));

            if (authors.Count > 0)
            {
                authorString += " and " + lastAuthor.ToString();
            }
            else
            {
                authorString = lastAuthor.ToString();
            }

            string titleString = Title != "" ? $"\"{Title}\"" : "[WARNING]DOES_NOT_FIND_TITLE";
            string informationString = Information.ToString();
            informationString = informationString != "" ? informationString : "[WARNING]DOES_NOT_FIND_INFORMATION";

            string result = string.Join(", ", [authorString, titleString, informationString]);
            return result;
        }

        public string Serialize()
        {
            return JsonSerializer.Serialize(this, new JsonSerializerOptions
            {
                WriteIndented = true,
                PropertyNamingPolicy = JsonNamingPolicy.SnakeCaseLower,
                Encoder = System.Text.Encodings.Web.JavaScriptEncoder.UnsafeRelaxedJsonEscaping,
            });
        }

        public static Paper? Deserialize(string json)
        {
            return JsonSerializer.Deserialize<Paper>(json, new JsonSerializerOptions
            {
                PropertyNamingPolicy = JsonNamingPolicy.SnakeCaseLower,
                Encoder = System.Text.Encodings.Web.JavaScriptEncoder.UnsafeRelaxedJsonEscaping,
            });
        }
    }


    public enum Languages
    {
        Unknown = 0,
        Korean = 1,
        English = 2,
    }

}