using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Security;
using System.Text;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Threading.Tasks;

namespace JournalListUpManager
{
    public class Paper
    {
        public List<Author> Authors { get; set; } = [];
        public string Title { get; set; } = "";
        public Information Information { get; set; } = Information.Empty;

        public string Abstract { get; set; } = "";
        public string[] Keywords { get; set; } = [];
        public string[] Tags { get; set; } = [];
        public string[] References { get; set; } = [];
        public string Note { get; set; } = "";
        public string FilePath { get; set; } = "";

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

    public struct Page
    {
        public int Start { get; set; }
        public int End { get; set; }

        [JsonIgnore]
        public int Length => End - Start + 1;

        public static Page Unknown => new Page
        {
            Start = -1,
            End = -1,
        };

        public override string ToString()
        {
            if (Start == -1 && End == -1)
                return "";
            else if (Length == 1)
                return $"pp. {Start.ToString()}";
            else
                return $"pp. {Start}-{End}";
        }
    }

    public struct Identifiers
    {
        public string DOI { get; set; }
        public string ISBN { get; set; }
        public string ISSN { get; set; }
        public string EISSN { get; set; }

        public static Identifiers Empty => new Identifiers
        {
            DOI = "",
            ISBN = "",
            ISSN = "",
            EISSN = ""
        };

        public override string ToString()
        {
            List<string> results = [];
            if (!string.IsNullOrEmpty(DOI))
                results.Add($"DOI: {DOI}");
            if (!string.IsNullOrEmpty(ISBN))
                results.Add($"ISBN: {ISBN}");
            if (!string.IsNullOrEmpty(ISSN))
                results.Add($"ISSN: {ISSN}");
            if (!string.IsNullOrEmpty(EISSN))
                results.Add($"EISSN: {EISSN}");

            return string.Join(", ", results);
        }
    }

    public struct JDate
    {
        private static readonly string[] MonthNames = new[]
        {
            "January", "February", "March", "April", "May", "June",
            "July", "August", "September", "October", "November", "December"
        };

        public int Year { get; set; }
        public int Month { get; set; }

        [JsonIgnore]
        public int MonthIndex => Month - 1;
        [JsonIgnore]
        public string MonthFullName => MonthNames[MonthIndex];
        [JsonIgnore]
        public string MonthShortName => MonthNames[MonthIndex].Substring(0, 3) + ".";

        public static JDate Unknown => new JDate
        {
            Year = -1,
            Month = -1,
        };

        public override string ToString()
        {
            if (Year == -1 && Month == -1)
                return "";
            else if (Month == -1)
                return Year.ToString();
            else
                return $"{MonthShortName} {Year}";
        }
    }

    public struct Author
    {
        public int ID { get; set; }
        public string Name { get; set; }
        public string Affiliation { get; set; }
        public string Email { get; set; }
        public AuthorTitle AuthorTitle { get; set; }

        public static Author Empty => new Author
        {
            ID = -1,
            Name = "Unknown",
            Affiliation = "Unknown",
            Email = "Unknown",
            AuthorTitle = AuthorTitle.Unknown
        };

        public override string ToString()
        {
            return $"{Name}";
        }

        public Author Sync(Author target)
        {
            target.ID = ID;
            return target;
        }
    }

    public struct Information
    {
        public string JournalName { get; set; }
        public int Volume { get; set; }
        public int Number { get; set; }
        public string Issue { get; set; }
        public Page Page { get; set; }
        public JDate Date { get; set; }
        public Identifiers Identifiers { get; set; }
        public DB DB { get; set; }

        public Languages Language { get; set; }
        public string JournalURL { get; set; }

        public static Information Empty => new Information
        {
            JournalName = "",
            Volume = -1,
            Number = -1,
            Issue = "",
            Page = Page.Unknown,
            Date = JDate.Unknown,
            Identifiers = Identifiers.Empty,
            DB = DB.None,
            Language = Languages.Unknown,
            JournalURL = ""
        };

        public override string ToString()
        {
            string result = "";
            List<string> results = [];

            if (!string.IsNullOrEmpty(JournalName))
                results.Add($"{JournalName}");
            if (Volume != -1)
                results.Add($"Vol. {Volume}");
            if (Number != -1)
                results.Add($"No. {Number}");
            if (!string.IsNullOrEmpty(Issue))
                results.Add($"Issue {Issue}");
            if (Page.ToString() != "")
                results.Add(Page.ToString());
            if (Date.ToString() != "")
                results.Add(Date.ToString());
            if (Identifiers.ToString() != "")
                results.Add(Identifiers.ToString());

            result = string.Join(", ", results);

            if (DB != DB.None)
            {
                List<string> dbResults = [];

                if (DB.HasFlag(DB.KCI))
                    dbResults.Add("KCI");
                if (DB.HasFlag(DB.SCOPUS))
                    dbResults.Add("SCOPUS");
                if (DB.HasFlag(DB.SCI))
                    dbResults.Add("SCI");
                if (DB.HasFlag(DB.SCIE))
                    dbResults.Add("SCIE");
                if (DB.HasFlag(DB.IEEE))
                    dbResults.Add("IEEE");

                result += $"({string.Join(", ", dbResults)})";
            }

            return result;
        }
    }

    public enum AuthorTitle
    {
        Unknown = 0,
        FirstAuthor = 1,
        CoAuthor = 2,
        CorrespondingAuthor = 3,
    }

    public enum Languages
    {
        Unknown = 0,
        Korean = 1,
        English = 2,
    }

    public enum DB
    {
        None = 0,
        KCI = 1,
        SCOPUS = 2,
        SCI = 4,
        SCIE = 8,
        IEEE = 16,
    }
}