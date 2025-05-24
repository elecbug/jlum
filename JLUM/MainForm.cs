using Microsoft.Web.WebView2.WinForms;

namespace JLUM
{
    public class MainForm : Form
    {
        public SplitContainer MainSplitContainer { get; private set; }

        /// <summary>
        /// MainForm constructor that initializes the form with the provided arguments.
        /// </summary>
        /// <param name="args"></param>
        public MainForm(string[] args)
        {
            Setup(args);

            // Set up the main form properties
            {
                Text = "Journal List-up Manager";
                Location = MainFormSettings.Default.Location;
                StartPosition = FormStartPosition.Manual;
                Size = MainFormSettings.Default.Size;
                Padding = new Padding(10);
                Font = new Font("Arial", 10);

                FormClosing += MainFormClosing;
            }

            GroupBox mainGroupBox = new GroupBox()
            {
                Parent = this,
                Dock = DockStyle.Fill,
                Text = MainFormSettings.Default.JournalListUpDirectory,
                Padding = new Padding(5),
            };

            MainSplitContainer = new SplitContainer()
            {
                Parent = mainGroupBox,
                Dock = DockStyle.Fill,
                Orientation = Orientation.Vertical,
                SplitterDistance  = MainFormSettings.Default.SplitterDistance,
            };

            WebView2 webView = new WebView2()
            {
                Parent = MainSplitContainer.Panel1,
                Dock = DockStyle.Fill,
                Source = new Uri(MainFormSettings.Default.LatestViewPDF),
            };
        }

        /// <summary>
        /// Sets up the main form based on the provided command line arguments.
        /// </summary>
        /// <param name="args"></param>
        private void Setup(string[] args)
        {
            if (args.Length > 1)
            {
                MessageBox.Show(
                    "This program only accepts one argument, which is the path to the journal list-up directory.",
                    "Error",
                    MessageBoxButtons.OK,
                    MessageBoxIcon.Error);
                Environment.Exit(1);
            }
            else if (args.Length == 1)
            {
                if (!Directory.Exists(args[0]))
                {
                    MessageBox.Show(
                        "The specified directory does not exist.",
                        "Error",
                        MessageBoxButtons.OK,
                        MessageBoxIcon.Error);
                    Environment.Exit(1);
                }
                else
                {
                    MainFormSettings.Default.JournalListUpDirectory = args[0];
                    MainFormSettings.Default.LatestViewPDF = args[0];
                    MainFormSettings.Default.Save();
                }
            }
            else if (args.Length == 0)
            {
                if (string.IsNullOrEmpty(MainFormSettings.Default.JournalListUpDirectory) ||
                    !Directory.Exists(MainFormSettings.Default.JournalListUpDirectory))
                {
                    MessageBox.Show(
                        "No journal list-up directory specified. Please provide a valid directory as an argument or set it in the settings.",
                        "Error",
                        MessageBoxButtons.OK,
                        MessageBoxIcon.Error);
                    Environment.Exit(1);
                }
            }
        }

        /// <summary>
        /// Handles the form closing event to prompt the user for confirmation before exiting the application.
        /// </summary>
        /// <param name="sender"></param>
        /// <param name="e"></param>
        private void MainFormClosing(object? sender, FormClosingEventArgs e)
        {
            DialogResult result = MessageBox.Show(
                "Exit the Program?",
                "Question",
                MessageBoxButtons.YesNo,
                MessageBoxIcon.Question);

            if (result == DialogResult.No)
            {
                e.Cancel = true;
            }
            else if (result == DialogResult.Yes)
            {
                MainFormSettings.Default.Location = Location;
                MainFormSettings.Default.Size = Size;
                MainFormSettings.Default.SplitterDistance = MainSplitContainer.SplitterDistance;

                MainFormSettings.Default.Save();
            }
        }
    }
}
