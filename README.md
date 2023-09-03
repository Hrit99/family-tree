# family-tree
github.com/Hrit99/family-tree

A cli command to track people relationship records.

It provides feature to : 

1. add a new person (name and gender). [People can have same names and logic assigns new ids to them. Used hashmap to store and search a person.] Gender is optional i.e. if not specified using tag, the person is treated as unknown gender and logic is implemented in that manner.
    family-tree add person <name> -g <gender(male/female)>
2. create a new relation (relation name, relation opposite to it i.e. father : child (son/daughter)). [New relations with new names can be created i.e. maternal-aunt. Relations persists and need not to be created again. ] Gender is optional i.e. if not specified using tag, the relation is treated as unknown gender and logic is implemented in that manner.
    family-tree add relation <relation> -g <gender(male/female)>
3. connect two person with a relation created. [Prompts to select person with id if more than one person with same name exists.]
    family-tree connect <name1> as <relation> of <name2>
4. queries :
    a. count people with particular relation to a person.
        family-tree count <relation> of <name>
    b. get person with particular relatioon to a person.
        family-tree get <relation> of <name>

IF USING SPACES IN NAMING PERSON OR RELATIONS THEN USE DOUBLE QUOTES i.e "SOME NAME"