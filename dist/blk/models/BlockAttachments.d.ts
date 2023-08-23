/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { AttachmentType } from './AttachmentType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Return object to the request for information regarding block attachments.
 * @export
 * @interface BlockAttachments
 */
export interface BlockAttachments {
    /**
     * Attachment List.
     * @type {Array<AttachmentType>}
     * @memberof BlockAttachments
     */
    blockAttachments?: Array<AttachmentType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof BlockAttachments
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof BlockAttachments
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the BlockAttachments interface.
 */
export declare function instanceOfBlockAttachments(value: object): boolean;
export declare function BlockAttachmentsFromJSON(json: any): BlockAttachments;
export declare function BlockAttachmentsFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockAttachments;
export declare function BlockAttachmentsToJSON(value?: BlockAttachments | null): any;
