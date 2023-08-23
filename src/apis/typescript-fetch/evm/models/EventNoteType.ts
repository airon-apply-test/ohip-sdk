/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Pertain event's comment information.
 * @export
 * @interface EventNoteType
 */
export interface EventNoteType {
    /**
     * 
     * @type {UniqueIDType}
     * @memberof EventNoteType
     */
    noteId?: UniqueIDType;
    /**
     * Note code.
     * @type {string}
     * @memberof EventNoteType
     */
    noteCode?: string;
    /**
     * Indicates if comment are internal use only.
     * @type {boolean}
     * @memberof EventNoteType
     */
    internal?: boolean;
    /**
     * Pertain comment tile.
     * @type {string}
     * @memberof EventNoteType
     */
    noteTitle?: string;
    /**
     * Pertain comment text.
     * @type {string}
     * @memberof EventNoteType
     */
    comment?: string;
    /**
     * Pertain display sequence.
     * @type {number}
     * @memberof EventNoteType
     */
    sequence?: number;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof EventNoteType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof EventNoteType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof EventNoteType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof EventNoteType
     */
    lastModifierId?: string;
    /**
     * Date an item will be purged from a database (e.g., from a live database to an archive).
     * @type {Date}
     * @memberof EventNoteType
     */
    purgeDate?: Date;
}

/**
 * Check if a given object implements the EventNoteType interface.
 */
export function instanceOfEventNoteType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function EventNoteTypeFromJSON(json: any): EventNoteType {
    return EventNoteTypeFromJSONTyped(json, false);
}

export function EventNoteTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EventNoteType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'noteId': !exists(json, 'noteId') ? undefined : UniqueIDTypeFromJSON(json['noteId']),
        'noteCode': !exists(json, 'noteCode') ? undefined : json['noteCode'],
        'internal': !exists(json, 'internal') ? undefined : json['internal'],
        'noteTitle': !exists(json, 'noteTitle') ? undefined : json['noteTitle'],
        'comment': !exists(json, 'comment') ? undefined : json['comment'],
        'sequence': !exists(json, 'sequence') ? undefined : json['sequence'],
        'createDateTime': !exists(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !exists(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !exists(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !exists(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !exists(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
    };
}

export function EventNoteTypeToJSON(value?: EventNoteType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'noteId': UniqueIDTypeToJSON(value.noteId),
        'noteCode': value.noteCode,
        'internal': value.internal,
        'noteTitle': value.noteTitle,
        'comment': value.comment,
        'sequence': value.sequence,
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0,10)),
    };
}

