/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { FormattedTextTextType } from './FormattedTextTextType';
import {
    FormattedTextTextTypeFromJSON,
    FormattedTextTextTypeFromJSONTyped,
    FormattedTextTextTypeToJSON,
} from './FormattedTextTextType';

/**
 * An indication of a new paragraph for a sub-section of a formatted text message.
 * @export
 * @interface CommentType
 */
export interface CommentType {
    /**
     * 
     * @type {FormattedTextTextType}
     * @memberof CommentType
     */
    text?: FormattedTextTextType;
    /**
     * An image for this paragraph.
     * @type {string}
     * @memberof CommentType
     */
    image?: string;
    /**
     * A URL for this paragraph.
     * @type {string}
     * @memberof CommentType
     */
    url?: string;
    /**
     * Specifies Comment's Title.
     * @type {string}
     * @memberof CommentType
     */
    commentTitle?: string;
    /**
     * Notification Location associated with the Note.
     * @type {string}
     * @memberof CommentType
     */
    notificationLocation?: string;
    /**
     * Specifies type of the comment.
     * @type {string}
     * @memberof CommentType
     */
    type?: string;
    /**
     * Comment type Description.
     * @type {string}
     * @memberof CommentType
     */
    typeDescription?: string;
    /**
     * When true, the comment may not be shown to the consumer. When false, the comment may be shown to the consumer.
     * @type {boolean}
     * @memberof CommentType
     */
    internal?: boolean;
    /**
     * When true, the comment may be confidential.
     * @type {boolean}
     * @memberof CommentType
     */
    confidential?: boolean;
    /**
     * When true, the note internal could be modified.
     * @type {boolean}
     * @memberof CommentType
     */
    overrideInternal?: boolean;
    /**
     * When true, the note title will be populated from the note type description and couldn't be modified.
     * @type {boolean}
     * @memberof CommentType
     */
    protectDescription?: boolean;
    /**
     * If specified comment belongs to the Hotel, otherwise it is a global comment.
     * @type {string}
     * @memberof CommentType
     */
    hotelId?: string;
    /**
     * Specifies type of action described in the comments.
     * @type {string}
     * @memberof CommentType
     */
    actionType?: string;
    /**
     * Indicates at which date an action described in the comment must be taken.
     * @type {Date}
     * @memberof CommentType
     */
    actionDate?: Date;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof CommentType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof CommentType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof CommentType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof CommentType
     */
    lastModifierId?: string;
    /**
     * Date an item will be purged from a database (e.g., from a live database to an archive).
     * @type {Date}
     * @memberof CommentType
     */
    purgeDate?: Date;
}

/**
 * Check if a given object implements the CommentType interface.
 */
export function instanceOfCommentType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CommentTypeFromJSON(json: any): CommentType {
    return CommentTypeFromJSONTyped(json, false);
}

export function CommentTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CommentType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'text': !exists(json, 'text') ? undefined : FormattedTextTextTypeFromJSON(json['text']),
        'image': !exists(json, 'image') ? undefined : json['image'],
        'url': !exists(json, 'url') ? undefined : json['url'],
        'commentTitle': !exists(json, 'commentTitle') ? undefined : json['commentTitle'],
        'notificationLocation': !exists(json, 'notificationLocation') ? undefined : json['notificationLocation'],
        'type': !exists(json, 'type') ? undefined : json['type'],
        'typeDescription': !exists(json, 'typeDescription') ? undefined : json['typeDescription'],
        'internal': !exists(json, 'internal') ? undefined : json['internal'],
        'confidential': !exists(json, 'confidential') ? undefined : json['confidential'],
        'overrideInternal': !exists(json, 'overrideInternal') ? undefined : json['overrideInternal'],
        'protectDescription': !exists(json, 'protectDescription') ? undefined : json['protectDescription'],
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'actionType': !exists(json, 'actionType') ? undefined : json['actionType'],
        'actionDate': !exists(json, 'actionDate') ? undefined : (new Date(json['actionDate'])),
        'createDateTime': !exists(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !exists(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !exists(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !exists(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !exists(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
    };
}

export function CommentTypeToJSON(value?: CommentType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'text': FormattedTextTextTypeToJSON(value.text),
        'image': value.image,
        'url': value.url,
        'commentTitle': value.commentTitle,
        'notificationLocation': value.notificationLocation,
        'type': value.type,
        'typeDescription': value.typeDescription,
        'internal': value.internal,
        'confidential': value.confidential,
        'overrideInternal': value.overrideInternal,
        'protectDescription': value.protectDescription,
        'hotelId': value.hotelId,
        'actionType': value.actionType,
        'actionDate': value.actionDate === undefined ? undefined : (value.actionDate.toISOString().substr(0,10)),
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0,10)),
    };
}

