DECLARE SUB where (d#, pants#)
DECLARE SUB look (d#)
DECLARE SUB lookat (d#)
DIM level#(4, 4)
CLS
FOR c# = 1 TO 4
FOR b# = 1 TO 4
level#(b#, c#) = a#
a# = a# + 1
NEXT
NEXT
PRINT "Save Mel Gibsons Testicles"
PRINT "By: Jesse Donat, Andrew Gross, Jeff Forshee"
PRINT "Testicles Provided by: Nick Miller"
PRINT ""
b# = 1: c# = 1
DO
d# = level#(b#, c#)
where d#, pants#
INPUT a$
a$ = LCASE$(a$)
IF a$ = "look" THEN look (d#)
IF a$ = "look at" THEN lookat (d#)
'WALK
IF a$ = "go east" THEN b# = b# + 1
IF a$ = "go west" THEN b# = b# - 1
IF a$ = "go south" THEN c# = c# + 1
IF a$ = "go north" THEN c# = c# - 1
IF a$ = "help" THEN PRINT "go, look, look at, climb, take, and many more"
'MISC
IF a$ = "take kilt" AND pants# = 0 AND d# = 4 THEN
        pants# = 1
        PRINT "huzzah, you have the kilt"
        END IF
IF a$ = "overthrow" THEN
        PRINT "You have overthrown england. as searching your new castle, in the treasure"
        PRINT "room you find a pair of Joe Nameth sling shot briefs"
        PRINT "*YOU HAVE RECIEVED SLING SHOT BRIEFS*"
        breifs# = 1
        END IF
IF a$ = "ride sombrero" THEN
        PRINT "In the process of riding the sombrero you lose your kilt and end up in Paris"
        pants# = 0
        b# = 4
        c# = 1
        IF breifs# = 0 THEN
                PRINT "The space police notice you standing nude in the center of Paris,"
                PRINT "rather than cause a sceen they choose to proceed to shoot you in the"
                PRINT "testicles. You die from loss of blood."
                GOTO 10
                END IF
        END IF
'CLIMB
IF a$ = "climb fence" AND pants# = 0 AND d# = 1 THEN
        PRINT "As you begin climbing the fence you feel a sharp pull."
        PRINT "You continue to climb the fence as a loud noice comes from your crotch,"
        PRINT "then another from your mouth."
        PRINT "As you look down you see your testicles have failed to make the climb with"
        PRINT "the rest of your body. They are hanging on the fence. "
        PRINT "You slowly die from an agonizing pain and a severe amount of blood loss."
        GOTO 10
        END IF
IF a$ = "climb fence" AND pants# = 1 AND d# = 1 THEN b# = b# + 1
'BOUNDRIES
IF a$ = "go east" THEN
        IF d# = 4 OR d# = 3 OR d# = 12 THEN
                PRINT "That is ocean, you fool"
                b# = b# - 1
                END IF
        END IF
IF a$ = "go west" THEN
        IF d# = 0 OR d# = 4 OR d# = 12 THEN
                PRINT "That is the ocean, you fool"
                b# = b# + 1
                END IF
        END IF
IF a$ = "go south" THEN
        IF d# = 4 OR d# = 1 OR d# = 2 OR d# = 3 OR d# = 12 THEN
                PRINT "That is the ocean, you fool"
                c# = c# - 1
                END IF
        END IF
IF a$ = "go north" THEN
        IF d# = 0 OR d# = 1 OR d# = 2 OR d# = 3 OR d# = 12 THEN
        PRINT "That is ocean, you fool"
        c# = c# + 1
        END IF
        END IF
IF a$ = "go east" AND d# = 1 THEN
        b# = b# - 1
        PRINT "their is a fence in the way, you must climb the fence"
        END IF
LOOP
'ENDINGS
10
DO
INPUT a$
PRINT "You are dead, you jack ass"
LOOP

SUB look (d#)
IF d# = 0 THEN PRINT "east: fence  north: ocean  south: corpse"
IF d# = 1 THEN PRINT "east: england  north: ocean  south: ocean  west: where you were nude"
IF d# = 2 THEN PRINT "east: sombrero  north: ocean  south: exotic pelvic dancing  west: fence"
IF d# = 4 THEN PRINT "east: ocean  west: ocean  north: where you were naked  south: ocean"

END SUB

SUB lookat (d#)
IF d# = 0 THEN PRINT "It is a field"
IF d# = 1 THEN PRINT "You see a fence on which you could easily catch your testicles"
IF d# = 2 THEN PRINT "You see england in need of overthrowing"
IF d# = 3 THEN PRINT "You see a giant sombrero in need of riding"
IF d# = 4 THEN
        PRINT "Corpse is wearing YOUR kilt! It is the theiving corpse! You must take back"
        PRINT "what is rightfully yours!"
        END IF
END SUB

SUB where (d#, pants#)
IF pants# = 0 THEN
        IF d# = 0 THEN PRINT "You are standing naked in a field"
        IF d# = 1 THEN PRINT "You are standing naked at fence"
        IF d# = 4 THEN PRINT "You are standing naked at a corpse"
        IF d# = 12 THEN PRINT "You are standing in breifs in Paris"
        END IF
IF pants# = 1 THEN
        IF d# = 0 THEN PRINT "You are wearing a kilt in a field"
        IF d# = 1 THEN PRINT "You are wearing a kilt at fence"
        IF d# = 2 THEN PRINT "You are wearing a kilt in england"
        IF d# = 3 THEN PRINT "You are wearing a kilt near a giant sombrero"
        IF d# = 4 THEN PRINT "You are wearing a kilt at a naked corpse"
        END IF
END SUB

