/****************************************************************************/
/*                                                                          */
/* adbdump.c                                                                */
/*                                                                          */
/* Read HP100LX .ADB file and convert it to ASCII format                    */
/*                                                                          */
/* A. Garzotto, April '94                                                   */
/* 2025/12: RSN604 Torio Modifyed for GCC and adjust format.                 */
/*                                                                          */
/****************************************************************************/

#include <stdio.h>
#include <fcntl.h>

#define YES 1
#define NO  0
#define NIL 0

#define INT(p, n) (((int)p[n] & 255) + (((int)p[n + 1] & 255) << 8))

#ifndef O_BINARY    /* MSDOS needs this - UNIX doesn't know it */
#define O_BINARY 0
#endif

#define S_ALARM     1
#define S_CHECKOFF  2
#define S_MONTHLY   2
#define S_CARRY     4
#define S_WEEKLY    4
#define S_TODO     16
#define S_EVENT    32
#define S_STUB     64
#define S_APPT    128

#define S_NOREPEAT 256
#define S_REPEAT   512
#define S_MAC     1024

/****************************************************************************/

struct rep_desc
{
   char freq;
   int days;
   int month;
   char y1;
   char m1;
   char d1;
   char y2;
   char m2;
   char d2;
   char ndeleted;
   char *deleted;
};

typedef struct rep_desc *REPEAT;

struct rec_desc
{
   char reptype;            /* repeat type */
   char state;              /* record state */
   char deleted;            /* this is a deleted record */
   char *desc;              /* description */
   char *location;          /* location */
   char *category;          /* category??? */
   char year;               /* start date; */
   char month;
   char day;
   int stime;               /* start time */
   int etime;               /* end time */
   int duration;            /* durations in days */
   char prio[2];            /* todo priority */
   int lead;                /* alarm lead time */
   long timestamp;          /* time used for sorting list */
   int notenum;             /* number of corresponding note */
   char *note;              /* note */
   REPEAT repeat;           /* pointer to repeat description */
   struct rec_desc *next;   /* pointer to next element in list */
};

typedef struct rec_desc *RECORD;

struct note_desc
{
   int num;
   char *note;
   struct note_desc *next;
};

typedef struct note_desc *NOTE;

/****************************************************************************/

//FILE *fout = stdout;     /* file to which the output is written */
FILE *fout;

int debug = NO;          /* debug mode */
char dateformat[80] = "d.m.yyyy"; /* date output format */
int cdf = NO;            /* output in comma delimited format */
int finished = NO;       /* set, if the lookup table is found (last record) */
int retain_crlf = NO;    /* retain CR/LF in CDF output */

char tquotes[2] = "\"";  /* quotes around CDF text fields */
char nquotes[2] = "";    /* quotes around CDF number fields */

int include = S_APPT|S_EVENT|S_TODO|S_NOREPEAT|S_REPEAT|S_MAC; 
                         /* record types included in output */
                         
                         
RECORD records = NIL;    /* list of data records */
NOTE notes = NIL;        /* list of note fields */

char *monthnames[] =
{
   "JAN",
   "FEB",
   "MAR",
   "APR",
   "MAY",
   "JUN",
   "JUL",
   "AUG",
   "SEP",
   "OCT",
   "NOV",
   "DEC"
};

char *monthnames1[] =
{
   "January",
   "February",
   "March",
   "April",
   "May",
   "June",
   "July",
   "August",
   "September",
   "October",
   "November",
   "December"
};

char *daynames[] =
{
   "Monday",
   "Tuesday",
   "Wednesday",
   "Thursday",
   "Friday",
   "Saturday",
   "Sunday",
   "?"
};

/****************************************************************************/
/* allocate memory */

char *my_malloc(size)
int size;
{
   char *p;
   
   p = (char *)malloc(size);
   if (!p)
   {
      fprintf(stderr, "Memory allocation problems.\nAborted!\n");
      exit(1);
   }
   return p;
}

/****************************************************************************/
/* create a new record struct */

RECORD new_rec()
{
   RECORD r;
   
   r = (RECORD)my_malloc(sizeof(struct rec_desc));
   r->reptype = 0;
   r->year = r->month = r->day = r->state = r->prio[0] = r->prio[1] = 0;
   r->stime = r->etime = r->duration = r->lead = 0;
   r->desc = r->location = r->category = r->note = NIL;
   r->timestamp = 0L;
   r->notenum = -1;
   r->repeat = NIL;
   r->next = NIL;
   r->deleted = NO;
   return r;
}

/****************************************************************************/
/* create a new repeat struct */

REPEAT new_rep()
{
   REPEAT r;
   
   r = (REPEAT)my_malloc(sizeof(struct rep_desc));
   r->freq = r->y1 = r->m1 = r->d1 = r->y2 = r->m2 = r->d2 = '\0';
   r->ndeleted = '\0';
   r->deleted = NIL;
   r->days = r->month = 0;
   return r;
}

/****************************************************************************/
/* insert record into ordered records list */

void rec_insert(r)
RECORD r;
{
   RECORD r1 = records, r2 = NIL;
   
   if (!r1)
   {
      records = r;
      return;
   }
   while (r1 && (r->timestamp > r1->timestamp))
   {
      r2 = r1;
      r1 = r1->next;
   }
   if (r1)
   {
      if (r2)
      {
         r->next = r2->next;
         r2->next = r;
      }
      else
      {
         r->next = records;
         records = r;
      }
   }
   else
      r2->next = r;
}

/****************************************************************************/
/* parse database header */

void dbheader(data)
char *data;
{
   if (data[2] != '2')
   {
      fprintf(stderr, "Not an .ADB file (%d)\n", data[2]);
      exit(1);
   }
   
   if (debug)
   {
      printf("HEADER\n");
      printf(" %d records\n", INT(data, 6));
   }
}

/****************************************************************************/
/* parse card layout definition */

void carddef(data)
char *data;
{
   if (debug) printf("CARD LAYOUT DEFINITION\n");
}

/****************************************************************************/
/* parse category list */

void category(data)
char *data;
{
   if (debug)
   {
      printf("CATEGORY LIST\n");
      printf(" %s\n", data);
   }
}

/****************************************************************************/
/* parse field definition */

void fielddef(data)
char *data;
{
   if (debug)
   {
      printf("FIELD DEFINITION (type %d, name '%s')\n",
             (int)data[0], &data[7]);
   }
}

/****************************************************************************/
/* parse viewpoint (sort and subset) definition */

void viewptdef(data)
char *data;
{
   if (debug) printf("VIEWPOINT DEFINITION\n");
}

/****************************************************************************/
/* parse note */

void note(data, len, index)
char *data;
int len, index;
{
   NOTE n;
   
   if (debug)
      printf("NOTE %d\n", index);
   
   data[len] = '\0';
   n = (NOTE)my_malloc(sizeof(struct note_desc));
   n->note = (char *)my_malloc(len + 1);
   strcpy(n->note, data);
   n->num = index;
   n->next = notes;
   notes = n;
}

/****************************************************************************/
/* parse table of viewpoint entries */

void viewpttable(data)
char *data;
{
   if (debug) printf("VIEWPOINT TABLE\n");
}

/****************************************************************************/
/* create a repeat struct and fill it with given data */

void create_repeat(data, rec)
char *data;
RECORD rec;
{
   REPEAT r;
   int i, len;
   
   r = new_rep();
   r->freq = data[0];
   r->days = INT(data, 1);
   r->month = INT(data, 3);
   r->y1 = data[5];
   r->m1 = data[6];
   r->d1 = data[7];
   r->y2 = data[8];
   r->m2 = data[9];
   r->d2 = data[10];
   r->ndeleted = data[11];
   if (r->ndeleted)
   {
      len = 4 * ((int)r->ndeleted & 255);
      r->deleted = (char *)my_malloc(len);
      for (i = 0; i < len; i++) r->deleted[i] = data[i + 12];
   }
   rec->repeat = r;
}

/****************************************************************************/
/* parse record data */

void recdata(data, state)
char *data;
char state;
{
   RECORD r = new_rec();
   char *p;
   
   if (debug) printf("DATA '%s'\n", &data[27]);
      
   r->state = data[14];
   
   if ((state != (char)2) && (state != (char)0)) r->deleted = (int)state;
   
   r->reptype = data[26];
   
   p = &data[27];
   r->desc = (char *)my_malloc(strlen(p) + 1);
   strcpy(r->desc, p);
   
   p = &data[INT(data, 4)];
   if (*p)
   {
      r->location = (char *)my_malloc(strlen(p) + 1);
      strcpy(r->location, p);
   }
   p = &data[INT(data, 2)];
   if (*p)
   {
      r->category = (char *)my_malloc(strlen(p) + 1);
      strcpy(r->category, p);
   }
   
   r->year = data[15];
   r->month = data[16];
   r->day = data[17];
   
   if (r->state & S_APPT)
   {
      r->stime = INT(data, 18);
      r->etime = INT(data, 22);
   }
   if (r->state & S_TODO)
   {
      r->prio[0] = data[18];
      r->prio[1] = data[19];
      if (!r->prio[0]) r->prio[0] = ' ';
      if (!r->prio[1]) r->prio[1] = ' ';
   }
   
   r->duration = INT(data, 20);
   
   if (r->state & S_APPT)
      r->lead = INT(data, 24);
   
   r->notenum = INT(data, 8);
   
   r->timestamp = (long)(r->stime) + 1440L * (long)(r->day) +
                  44640L * (long)(r->month) + 535680L * (long)(r->year);
   
   if (r->reptype != (char)1)
      create_repeat(&data[INT(data, 6)], r);
   rec_insert(r);
}

/****************************************************************************/
/* parse smart clip record */

void linkdef(data)
char *data;
{
   if (debug) printf("SMART CLIP (%s)\n", &data[4]);
}

/****************************************************************************/
/* parse multiple card definition */

void cardpagedef(data)
char *data;
{
   if (debug) printf("MULTIPLE CARD DEFINITION\n");
}

/****************************************************************************/
/* parse lookup table */

void lookuptable(data)
char *data;
{
   if (debug) printf("LOOKUP TABLE\n");
   
   finished = YES;
}

/****************************************************************************/
/* parse appt book info table */

void appt_info(data)
char *data;
{
   if (debug) printf("APPT BOOK INFO\n");
}

/****************************************************************************/
/* parse appt book list record */

void appt_list(data, len)
char *data;
int len;
{
   if (debug) printf("APPT BOOK LIST\n");
}

/****************************************************************************/
/* convert data from .ADB file into ASCII and print it to stdout */

void convert_adb(handle)
int handle;
{
   char buf[6];
   int len, index;
   char *data;
   
   read(handle, buf, 4);
   if ((buf[0] != 'h') || (buf[1] != 'c') ||
       (buf[2] != 'D') || (buf[3] != '\0'))
   {
      fprintf(stderr, "Not a HP100LX PIM file\n");
      exit(1);
   }
   
   while (!finished && (read(handle, buf, 6) == 6))
   {
      len = INT(buf, 2) - 6;
      index = INT(buf, 4);
      if (debug)
         printf("Type %d length %d index %d\n", (int)*buf, len, index);
      data = (char *)my_malloc(len + 2);
      if (read(handle, data, len) != len)
      {
         perror("Read error");
         exit(1);
      }
      switch(*buf)
      {
      case 0:  dbheader(data); break;
      case 4:  carddef(data); break;
      case 5:  category(data); break;
      case 6:  fielddef(data); break;
      case 7:  viewptdef(data); break;
      case 9:  note(data, len, index); break;
      case 10: viewpttable(data); break;
      case 11: recdata(data, buf[1]); break;
      case 12: linkdef(data); break;
      case 13: cardpagedef(data); break;
      case 14: appt_info(data); break;
      case 15: appt_list(data, len); break;
      case 31: lookuptable(data); break;
      default: fprintf(stderr, "Unknown header type: %d\n\n", (int)*buf);
               fprintf(stderr, "This file is probably password protected.\n");
               fprintf(stderr, "ADBDUMP can only process unprotected files.\n");
               exit(1);
      }
      free(data);
   }
}

/****************************************************************************/
/* place notes into corresponding data records */

void fixup_notes()
{
   NOTE n = notes;
   RECORD r;
   
   while (n)
   {
      r = records;
      while (r && (r->notenum != n->num)) r = r->next;
      if (r)
         r->note = n->note;
      else
         fprintf(stderr, "WARNING: Unreferenced note:\n%s\n", n->note);
      n = n->next;
   }
   
   r = records;
   while (r)
   {
      if ((r->notenum & 65535 != 65535) && !r->note)
         fprintf(stderr, "WARNING: Referenced note does not exist:\n%s\n",
                 r->desc);
      r = r->next;
   }
}

/****************************************************************************/
/* output text field */

void put_text(str)
char *str;
{
   char *p = str;
   
   fprintf(fout, "%s", tquotes);
   if (!p)
   {
      fprintf(fout, "%s", tquotes);
      return;
   }
   
   while (*p)
   {
      switch (*p)
      {
      case '\r': if (retain_crlf)
                    fprintf(fout, "\r");
                 else
                    fprintf(fout, "\\r");
                 break;
      case '\n': if (retain_crlf)
                    fprintf(fout, "\n");
                 else
                    fprintf(fout, "\\n");
                 break;
      case '\"': fprintf(fout, "\\\"");
                 break;
      case ',':  if (!*tquotes)
                    fprintf(fout, "\\,");
                 else
                    fprintf(fout, ",");
                 break;
      default: fprintf(fout, "%c", *p);
      }
      p++;
   }
   fprintf(fout, "%s", tquotes);
}

/****************************************************************************/
/* output date in the format given in 'dateformat' */

void put_date(y, m, d)
int y, m, d;
{
   char *p = dateformat;
   int i;
   
   while (*p)
   {
      switch (*p)
      {
      case 'd': i = 0;
                while (*p == 'd') {p++; i++; }
                if (i > 1)
                   fprintf(fout, "%.2d", d + 1);
                else
                   fprintf(fout, "%d", d + 1);
                break; 
      case 'm': i = 0;
                while (*p == 'm') { p++; i++; }
                if (i > 3)
                   fprintf(fout, "%s", monthnames1[m]);
                else if (i > 2)
                   fprintf(fout, "%s", monthnames[m]);
                else if (i > 1)
                   fprintf(fout, "%.2d", m + 1);
                else
                   fprintf(fout, "%d", m + 1);
                break; 
      case 'y': i = 0;
                while (*p == 'y') { p++; i++; }
                if (i > 2)
                   fprintf(fout, "%d", 1900 + (y & 255));
                else
                   fprintf(fout, "%d", (y & 255) % 100);
                break; 
      default: fprintf(fout, "%c", *p);
               p++;
               break;
      }
   }
}

/****************************************************************************/
/* output database in comma delimited format */

void output_cdf(r)
RECORD r;
{
   char *p;
   int i;
   
   fprintf(fout, "%s", tquotes);
   put_date(r->year, r->month, r->day);
   fprintf(fout, "%s,", tquotes);
   put_text(r->desc);
   fprintf(fout, ",");
   put_text(r->location);
   fprintf(fout, ",");
        
   /* do not output checked off to-dos; uncheck them! */ 
   if ((r->state & S_TODO) && (r->state & S_CHECKOFF))
      r->state ^= S_CHECKOFF;
      
   fprintf(fout, "%s%d%s,", nquotes, (int)r->state & 255, nquotes);
   if (r->state & S_APPT)
   {
      fprintf(fout, "%s%d:%.2d%s,", tquotes, r->stime / 60,
              r->stime % 60, tquotes);
      fprintf(fout, "%s%d:%.2d%s,", tquotes, r->etime / 60,
              r->etime % 60, tquotes);
      fprintf(fout, "%s%d%s,", nquotes, r->duration, nquotes);
      fprintf(fout, "%s%d%s,", nquotes, r->lead, nquotes);
   }
   else if (r->state & S_EVENT)
   {
      fprintf(fout, "%s%d%s,", nquotes, r->duration, nquotes);
   }
   else /* TODO */
   {
      fprintf(fout, "%s%c%c%s,", tquotes, r->prio[0], r->prio[1], tquotes);
      fprintf(fout, "%s%d%s,", nquotes, r->duration, nquotes);

	  // @@@@ RSN604 added to adjust format.  
      fprintf(fout, "%s%d%s,", nquotes, 0, nquotes);
      fprintf(fout, "%s%d%s,", nquotes, 0, nquotes);
	  // @@@@
   }
   
   put_text(r->note);
         
   if (r->repeat)
   {
      fprintf(fout, ",%s%d%s,", nquotes, (int)r->reptype & 255, nquotes);
      fprintf(fout, "%s%d%s,", nquotes, (int)r->repeat->freq & 255, nquotes);
      fprintf(fout, "%s%d%s,", nquotes, r->repeat->days, nquotes);
      fprintf(fout, "%s%d%s,", nquotes, r->repeat->month, nquotes);
      fprintf(fout, "%s%d%s,", nquotes, (int)r->repeat->y1 & 255, nquotes);
      fprintf(fout, "%s%d%s,", nquotes, (int)r->repeat->m1 & 255, nquotes);
      fprintf(fout, "%s%d%s,", nquotes, (int)r->repeat->d1 & 255, nquotes);
      fprintf(fout, "%s%d%s,", nquotes, (int)r->repeat->y2 & 255, nquotes);
      fprintf(fout, "%s%d%s,", nquotes, (int)r->repeat->m2 & 255, nquotes);
      fprintf(fout, "%s%d%s,", nquotes, (int)r->repeat->d2 & 255, nquotes);
      fprintf(fout, "%s%d%s", nquotes, (int)r->repeat->ndeleted&255,nquotes);
      for (i = 0; i < 4 * (int)r->repeat->ndeleted & 255; i++)
         fprintf(fout, ",%s%d%s", nquotes,
         (int)r->repeat->deleted[i] & 255, nquotes);
   }
   fprintf(fout, "\n");
}

/****************************************************************************/
/* output the date range in human readable format */

void output_range(r)
REPEAT r;
{
   int i, n;
   
   fprintf(fout, " (from ");
   put_date(r->y1, r->m1, r->d1);
   fprintf(fout, " to ");
   put_date(r->y2, r->m2, r->d2);
   fprintf(fout, ")\n");
   n = (int)(r->ndeleted) & 255;
   if (n)
   {
      fprintf(fout, " (except");
      for (i = 0; i < n; i += 4)
      {
         fprintf(fout, " ");
         put_date(r->deleted[i], r->deleted[i + 1], r->deleted[i + 2]);
      }
      fprintf(fout, ")\n");
   }
}

/****************************************************************************/
/* output day of week */

void output_weekday(days)
int days;
{
   int i, first = YES;
   
   for (i = 0; i < 7; i++)
   {
      if (days & (1 << i))
      {
         if (!first)
            fprintf(fout, ",");
         fprintf(fout, " %s", daynames[i]);
         first = NO;
      }
   }
}

/****************************************************************************/
/* output day field */

void output_days(days)
int days;
{
   
   if (days & 128)
   {
      fprintf(fout, " on the");
      if (days & 256)
         fprintf(fout, " 1st");
      if (days & 512)
         fprintf(fout, " 2nd");
      if (days & 1024)
         fprintf(fout, " 3rd");
      if (days & 2048)
         fprintf(fout, " 4th");
      if (days & 4096)
         fprintf(fout, " last");
      output_weekday(days);
   }
   else
   {
      fprintf(fout, " on the %d.", days);
   }
}

/****************************************************************************/
/* output month field */

void output_month(month)
int month;
{
   int i, first = YES;
   
   fprintf(fout, " of");
   for (i = 0; i < 12; i++)
   {
      if (month & (1 << i))
      {
         if (!first)
            fprintf(fout, ",");
         fprintf(fout, " %s", monthnames1[i]);
         first = NO;
      }
   }
}

/****************************************************************************/
/* output repeat type in human readable format */

void output_repeat(r_type, r)
int r_type;
REPEAT r;
{
   switch (r_type)
   {
   case 1:  /* not repeated */
            break; 
   case 2:  /* daily */
            if (r->freq > 1)
               fprintf(fout, "Repeated every %d days", (int)(r->freq)); 
            else
               fprintf(fout, "Repeated every day");
            output_range(r);
            break;
   case 4:  /* weekly */
            if (r->freq > 1)
               fprintf(fout, "Repeated weekly every %d.", (int)(r->freq)); 
            else
               fprintf(fout, "Repeated weekly every");
            output_weekday(r->days);
            output_range(r);
            break;
            
   case 8:  /* monthly */
            if (r->freq > 1)
               fprintf(fout, "Repeated every %d months", (int)(r->freq)); 
            else
               fprintf(fout, "Repeated monthly");
            output_days(r->days);
            output_range(r);
            break;
   case 16: /* yearly */
            if (r->freq > 1)
               fprintf(fout, "Repeated every %d years", (int)(r->freq)); 
            else
               fprintf(fout, "Repeated yearly");
            output_days(r->days);
            output_month(r->month);
            output_range(r);
           break;
   case 32: /* customized */
            fprintf(fout, "Repeated");
            output_days(r->days);
            output_month(r->month);
            output_range(r);
           break;
   default: fprintf(stderr, "WARNING: unknown repeat type %d!\n", r_type);
   }
}

/****************************************************************************/
/* output database in human readable format */

void output_readable(r)
RECORD r;
{
   fprintf(fout, "\n-------- %s --------\n",
      ((r->state & S_APPT)? "APPT":((r->state & S_EVENT)?
       "EVENT":"TODO")));
   put_date(r->year, r->month, r->day);
   if (r->duration && !(r->state & S_TODO))
      fprintf(fout,  " (next %d days)", r->duration);
   if (r->state & S_APPT)
      fprintf(fout, "  %d:%.2d to %d:%.2d",
              (r->stime) / 60,(r->stime) % 60,
              (r->etime) / 60, (r->etime) % 60);
   fprintf(fout, "\n");
   output_repeat((int)(r->reptype), r->repeat);
   if (r->duration && (r->state & S_TODO))
      fprintf(fout, "Due after %d days\n", r->duration - 1);
   if (r->state & S_TODO)
      fprintf(fout, "Priority: %c%c\n",
         r->prio[0], r->prio[1]);
   if ((r->state & S_CHECKOFF) && (r->state & S_TODO))
      fprintf(fout, "Checked off\n");
   fprintf(fout, "%s\n", r->desc);
   if (r->location)
      fprintf(fout, "@%s\n", r->location);
   if (r->category)
      fprintf(fout, "Category: %s\n", r->category);
   if (r->note)
      fprintf(fout, "\n%s\n", r->note);
}

/****************************************************************************/
/* print database in ASCII form */

void output_ascii()
{
   RECORD r = records;
   int doit;

   while (r)
   {
      doit = YES;
      if (!(include & S_NOREPEAT) && (r->reptype == (char)1)) doit = NO;
      if (!(include & S_REPEAT) && (r->reptype != (char)1)) doit = NO;
      if (!(include & (int)r->state & 255)) doit = NO;
      if (!(include & S_MAC) && (r->desc[0] == '|')) doit = NO;
      if (((int)r->state & 255 & S_APPT) && (include & S_MAC) &&
          (r->desc[0] == '|') &&
          (((include & S_NOREPEAT) && (r->reptype == (char)1)) ||
          (((include & S_REPEAT) && (r->reptype != (char)1)))))
           doit = YES;
      if (r->deleted)
      {
         fprintf(stderr, "WARNING: ignoring garbage record '%s' (%d)\n",
               r->desc, r->deleted);
         doit = NO;
      }
      if (doit)
      {
         if (cdf)
            output_cdf(r);
         else
            output_readable(r);
      }
      r = r->next;
   }
}

/****************************************************************************/
/* display usage information */

void help(prog)
char *prog;
{
   fprintf(stderr, "ADBDUMP version 1.3 by Andreas Garzotto\n\n");
   fprintf(stderr, "USAGE: %s [options] <.ADB file> [<outputfile>]\n", prog);
   fprintf(stderr, " Options:\n");
   fprintf(stderr, "  -c     output comma delimited format (CDF)\n");
   fprintf(stderr,
      "  -d fmt output date as specified in fmt (default: 'd.m.yyyy')\n");
   fprintf(stderr,
      "  -i xxx output includes the types xxx (default: 'aetmnr')\n");
   fprintf(stderr, "         (a=appointments e=events t=to-dos)\n");
   fprintf(stderr,
      "         (m=macros/DOS programs n=non-repeated r=repeated)\n");
   fprintf(stderr, "  -q0    do not put quotes around any fields in CDF\n");
   fprintf(stderr, "  -q1    put quotes around text only (default)\n");
   fprintf(stderr, "  -q2    put quotes around text and numbers\n");
   fprintf(stderr,
      "  -r     retain CR/LF in CDF (do not replace by \"\\r\\n\")\n");
   
   fprintf(stderr, "\n Example: %s -i en -c test.adb test.cdf\n", prog);
   fprintf(stderr,
      "   (writes non-repeating events from test.adb in comma\n");
   fprintf(stderr, "   delimited format to file test.cdf.)\n");

   exit(1);
}

/****************************************************************************/
/* decode options */

int get_options(argc, argv)
int argc;
char **argv;
{
   int handle = -1, i = 1;
   char *p;
   
   while (i < argc)
   {
      if (argv[i][0] == '-')
      {
         switch (argv[i][1])
         {
         case 'x': debug = YES; break;
         case 'd': if (i >= argc - 1) break;
                   strcpy(dateformat, argv[++i]);
                   break;
         case 'c': cdf = YES; break;
         case 'i': if (i >= argc - 1) break;
                   include = 0;
                   p = argv[++i];
                   while (*p)
                   {
                      switch (*p)
                      {
                      case 'a': include |= S_APPT; break;
                      case 'e': include |= S_EVENT; break;
                      case 't': include |= S_TODO; break;
                      case 'n': include |= S_NOREPEAT; break;
                      case 'r': include |= S_REPEAT; break;
                      case 'm': include |= S_MAC; break;
                      case 's': include |= S_STUB; break;
                      default: 
                         fprintf(stderr, 
                         "\nUnknown record type in -i option:'%c'.\n\n", *p);
                       help(argv[0]);
                      }
                      p++;
                   }
                   if (!(include & S_NOREPEAT) && !(include & S_REPEAT))
                      fprintf(stderr,
                      "WARNING: neither n nor r is set in -i option\n");
                   break;
         case 'q': if (argv[i][2] == '0') strcpy(tquotes, "");
                   if (argv[i][2] == '2') strcpy(nquotes, "\"");
                   break;
         case 'r': retain_crlf = YES; break;
         default: help(argv[0]);
         }
         if (retain_crlf && !*tquotes)
         {
            fprintf(stderr, "WARNING: multi line notes are not parsable\n");
            fprintf(stderr, "when using -r and -q0 at the same time!\n");
         }
      }
      else
      {
         if (handle >= 0)
            fout = fopen(argv[i], "w");
         else
         {
            handle = open(argv[i], O_RDONLY|O_BINARY);
            if (handle == -1)
            {
               fprintf(stderr, "Cannot open input file '%s'\n", argv[i]);
               exit(1);
            }
         }
      }
      i++;
   }
   return handle;
}

/****************************************************************************/

void main(argc, argv)
int argc;
char **argv;
{
   int handle = -1;
   fout = stdout;     /* file to which the output is written */   
   if (argc < 2) help(argv[0]);
   
   handle = get_options(argc, argv);
   if (handle == -1)
      help(argv[0]);
   convert_adb(handle);
   close(handle);
   fixup_notes();
   output_ascii();
   fclose(fout);
   exit(0);
}

/****************************************************************************/
